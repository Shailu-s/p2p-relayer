package p2p

import (
	"context"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-libp2p/p2p/discovery/util"
	"github.com/libp2p/go-libp2p/p2p/protocol/ping"
)

func (n *Node) sendPing(ctx context.Context, p peer.ID) error {
	pinger := ping.NewPingService(n.host)
	resCh := pinger.Ping(ctx, p)

	result := <-resCh
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func Discover(n *Node, interval time.Duration) {
	var routingDiscovery = routing.NewRoutingDiscovery(n.dht)

	util.Advertise(n.ctx, routingDiscovery, n.discoveryTag)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			peers, err := util.FindPeers(n.ctx, routingDiscovery, n.discoveryTag)
			if err != nil {
				n.logger.Warnf("failed to collect peers from discoverer: %v", err)
				continue
			}

			for _, p := range peers {
				if p.ID == n.host.ID() {
					continue
				}

				n.logger.Debugf("peer discovered: %s", p.String())

				connectedness := n.host.Network().Connectedness(p.ID)
				if connectedness == network.Connected {
					n.logger.Debugf("Connected (%s): %s", connectedness.String(), p.ID.String())
					continue
				}

				n.logger.Debugf("connecting to peer: %v", p.String())
				err = n.host.Connect(n.ctx, p)
				if err != nil {
					n.logger.Warnf("connect to peer %s failed: %v", p.ID.String(), err)
					continue
				}

				n.logger.Infof("connected to peer: %v", p.ID.String())
			}
		case <-n.ctx.Done():
			return
		}
	}
}

// discoveryNotifee gets notified when we find a new peer via mDNS discovery
type discoveryNotifee struct {
	h host.Host
	n Node
}

// setupDiscovery creates an mDNS discovery service and attaches it to the libp2p Host.
// This lets us automatically discover peers on the same LAN and connect to them.
func (n *Node) setupDiscovery(h host.Host, discoveryTag string) error {
	// setup mDNS discovery to find local peers
	s := mdns.NewMdnsService(h, discoveryTag, &discoveryNotifee{h: h, n: *n})
	return s.Start()
}

// HandlePeerFound connects to peers discovered via mDNS. Once they're connected,
// the PubSub system will automatically start interacting with them if they also
// support PubSub.
func (d *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	d.n.logger.Infof("discovered new peer %s\n", pi.ID.Pretty())
	err := d.h.Connect(context.Background(), pi)
	if err != nil {
		fmt.Printf("error connecting to peer %s: %s\n", pi.ID.Pretty(), err)
	}
}
