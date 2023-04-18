package p2p

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/sirupsen/logrus"
	protocols_p2p "github.com/volmexfinance/relayers/relayer-srv/chat"
)

// node client version
// TODO: add this in config
const clientVersion = "perp-relayer/0.0.1"
const DiscoveryInterval = 20 * time.Second

type Config struct {
	Port           int64
	ListenAddress  string
	DiscoveryTag   string
	RelayerAddress []string
}

type GuardianInfo struct {
	Address string  `json:"pubkey"`
	PeerId  peer.ID `json:"peer_id"`
}
type GuardianConfig struct {
	GuardianSet []GuardianInfo
}

// Node type - a p2p host implementing one or more p2p protocols
type Node struct {
	lock               sync.Mutex
	logger             *logrus.Entry
	host               host.Host // lib-p2p host
	dht                *dht.IpfsDHT
	chatRoom           *ChatRoom
	privKey            string
	ctx                context.Context
	discoveryTag       string
	GuardianSetMapping map[string][]GuardianInfo
}

// toAddrInfo returns discovered peers
func toAddrInfo(addresses []string) []*peer.AddrInfo {
	addrs := make([]*peer.AddrInfo, 0, len(addresses))
	for _, addr := range addresses {
		multiaddr, err := ma.NewMultiaddr(addr)
		if err != nil {
			panic(err)
		}

		info, err := peer.AddrInfoFromP2pAddr(multiaddr)
		if err != nil {
			panic(err)
		}
		addrs = append(addrs, info)
	}

	return addrs
}

func NewP2PNode(ctx context.Context, logger *logrus.Logger, obsRecieveRes chan<- *protocols_p2p.GossipMessage, obsSendReq <-chan *protocols_p2p.GossipMessage, privKey string, p2pCfg Config, guardianSetMapping map[string][]GuardianInfo) *Node {
	log := logger.WithField("layer", "p2p-node")
	bootstrapPeers := toAddrInfo(p2pCfg.RelayerAddress)

	log.Debugf("Lenth of relayer address %d", len(bootstrapPeers))

	h1 := makeRandomNode(p2pCfg, privKey, log, bootstrapPeers, ctx, obsRecieveRes, obsSendReq, guardianSetMapping)

	return h1
}

// helper method - create a lib-p2p host to listen on a port
func makeRandomNode(p2pCfg Config, privKeyString string, log *logrus.Entry, bootstrapNodes []*peer.AddrInfo, ctx context.Context, obsRecieveRes chan<- *protocols_p2p.GossipMessage, obsSendReq <-chan *protocols_p2p.GossipMessage, guardianSetMapping map[string][]GuardianInfo) *Node {
	priv, err := crypto.UnmarshalSecp256k1PrivateKey(common.Hex2Bytes(privKeyString))
	if err != nil {
		log.Panicf("cannot get p2p private key %s", err.Error())
	}

	listen, err := ma.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%d", p2pCfg.ListenAddress, p2pCfg.Port))
	if err != nil {
		log.Panicf("cannot create new multiaddress: %v", err)
	}

	nodeHost, err := libp2p.New(
		// Multiple listen addresses
		//TODO
		// libp2p.ListenAddrStrings(
		// 	// Listen on QUIC only.
		// 	// https://github.com/libp2p/go-libp2p/issues/688
		// 	fmt.Sprintf("/ip4/0.0.0.0/udp/%d/quic", port),
		// 	fmt.Sprintf("/ip6/::/udp/%d/quic", port),
		// ),
		libp2p.ListenAddrs(listen),

		libp2p.Identity(priv),
		libp2p.Ping(true),
	)
	if err != nil {
		log.Panicf("cannot initiate p2p node: %v", err)
	}

	// Add bootstrap nodes to host's peer store
	for _, bootstrapNode := range bootstrapNodes {
		nodeHost.Peerstore().AddAddrs(bootstrapNode.ID, bootstrapNode.Addrs, peerstore.PermanentAddrTTL)
	}

	// Initiate DHT to enable peer discovery
	nodeDht, err := NewDHT(ctx, log, nodeHost, bootstrapNodes)
	if err != nil {
		log.Panicf("cannot initiate nodeDht: %v", err)
	}

	// Initiate gossipsub
	gossipSub, err := pubsub.NewGossipSub(ctx, nodeHost)
	if err != nil {
		log.Panicf("failed to create a gossipusb instance: %v", err)
	}

	node := &Node{
		ctx:                ctx,
		host:               nodeHost,
		logger:             log,
		dht:                nodeDht,
		privKey:            privKeyString,
		discoveryTag:       p2pCfg.DiscoveryTag,
		GuardianSetMapping: guardianSetMapping,
	}

	// setup mDNS discovery
	if err := node.setupDiscovery(nodeHost, p2pCfg.DiscoveryTag); err != nil {
		log.Panicf("cannot setup mDSN discovery: %v", err)
	}

	chatRoom, err := JoinChatRoom(ctx, node, gossipSub, nodeHost.ID(), obsRecieveRes, obsSendReq)
	if err != nil {
		log.Panicf("cannot join chat room: %v", err)
	}
	node.chatRoom = chatRoom

	log.Infof("New node: ID: %s, port: %d\n", node.host.ID(), p2pCfg.Port)

	return node
}

func NewDHT(ctx context.Context, log *logrus.Entry, host host.Host, bootstrapPeers []*peer.AddrInfo) (*dht.IpfsDHT, error) {
	kdht, err := dht.New(ctx, host, dht.Mode(dht.ModeServer))
	if err != nil {
		return nil, err
	}

	if err = kdht.Bootstrap(ctx); err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	for _, peerInfo := range bootstrapPeers {
		wg.Add(1)
		peerInfo := peerInfo // copy to avoid race conditions
		go func() {
			defer wg.Done()
			if err := host.Connect(ctx, *peerInfo); err != nil {
				log.Infof("Error while connecting to node %q: %-v", *peerInfo, err)
				return
			}

			log.Infof("Connection established with bootstrap node: %q", *peerInfo)
		}()
	}
	wg.Wait()

	return kdht, nil
}

func (n *Node) Stop() {
	n.lock.Lock()
	defer n.lock.Unlock()
	n.logger.Infof("Stop signal received")
	n.chatRoom.sub.Cancel()
	if err := n.host.ConnManager().Close(); err != nil {
		panic(err)
	}
	n.logger.Infof("Closed ConnManager")
	if err := n.host.Network().Close(); err != nil {
		panic(err)
	}
	n.logger.Infof("Closed Network")
	if err := n.host.Peerstore().Close(); err != nil {
		panic(err)
	}
	n.logger.Infof("Closed Peerstore")
	if err := n.host.Close(); err != nil {
		panic(err)
	}
	n.logger.Infof("Closed host")
}

func (n *Node) AddToInputQueue() bool {
	if n.host.Peerstore().Peers().Len() > 0 {
		return true
	} else {
		n.logger.Info("Only one peer available")
		return false
	}
}

func (n *Node) GetPeerID() peer.ID {
	return n.host.ID()
}

func (n *Node) Run() {
	go Discover(n, DiscoveryInterval)
	go n.chatRoom.PublishMessageLoop()
	go n.chatRoom.ReadMessageLoop()
}

func (n *Node) AddNewOwner(owner common.Address, chain string) {
	// guardianSet := n.guardianSetMapping[chain]
	n.GuardianSetMapping[chain] = append(n.GuardianSetMapping[chain], GuardianInfo{Address: string(owner.Hex())})

	// n.guardianSetMapping = append(n.guardianSetMapping[chain], models.GuardianInfo{Address: string(owner.Hex())})
}

func (n *Node) RemoveOwner(owner common.Address, chain string) {
	for idx, v := range n.GuardianSetMapping[chain] {
		if v.Address == owner.Hex() {
			n.host.Network().ClosePeer(v.PeerId)
			n.GuardianSetMapping[chain] = append(n.GuardianSetMapping[chain][0:idx], n.GuardianSetMapping[chain][idx+1:]...)
		}
	}
}

func (n *Node) CheckStatus(peerID peer.ID) bool {
	return n.host.Network().Connectedness(peerID) == network.Connected
}
