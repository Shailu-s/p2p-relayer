package p2p

import (
	"context"
	"sync"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
	protocols_p2p "github.com/volmexfinance/relayers/relayer-srv/chat"
	"google.golang.org/protobuf/proto"
)

// ChatRoomBufSize is the number of incoming messages to buffer for each topic.
const ChatRoomBufSize = 128

var lock = sync.Mutex{}

// ChatRoom represents a subscription to a single PubSub topic. Messages
// can be published to the topic with ChatRoom.Publish, and received
// messages are pushed to the Messages channel.
type ChatRoom struct {
	// obsRecieveRes is a channel of messages received from other peers in the chat room
	obsRecieveRes chan<- *protocols_p2p.GossipMessage
	obsSendReq    <-chan *protocols_p2p.GossipMessage

	node  *Node
	ctx   context.Context
	ps    *pubsub.PubSub
	topic *pubsub.Topic
	sub   *pubsub.Subscription

	roomName string
	self     peer.ID
}

// JoinChatRoom tries to subscribe to the PubSub topic for the room name, returning
// a ChatRoom on success.
func JoinChatRoom(ctx context.Context, node *Node, ps *pubsub.PubSub, selfID peer.ID, obsRecieveRes chan<- *protocols_p2p.GossipMessage, obsSendReq <-chan *protocols_p2p.GossipMessage) (*ChatRoom, error) {
	// join the pubsub topic
	topic, err := ps.Join(topicName(node.discoveryTag))
	if err != nil {
		return nil, err
	}

	// and subscribe to it
	sub, err := topic.Subscribe()
	if err != nil {
		return nil, err
	}
	defer topic.Close()
	cr := &ChatRoom{
		ctx:           ctx,
		node:          node,
		ps:            ps,
		topic:         topic,
		sub:           sub,
		self:          selfID,
		roomName:      node.discoveryTag,
		obsRecieveRes: obsRecieveRes,
		obsSendReq:    obsSendReq,
	}

	// start reading messages from the subscription in a loop
	return cr, nil
}

// Publish sends a message to the pubsub topic.
func (cr *ChatRoom) Publish(msg *protocols_p2p.GossipMessage) error {
	lock.Lock()
	defer lock.Unlock()
	switch msg.Message.(type) {
	case *protocols_p2p.GossipMessage_MatchedOrder:
		msgData := msg.GetMatchedOrder().MessageData
		sig, err := cr.node.signData(msgData.Data)
		if err != nil {
			cr.node.logger.Errorf("sign error: %v", err)
		}
		msgData.Sign = append(msgData.Sign, sig)
		msg = &protocols_p2p.GossipMessage{
			Message: &protocols_p2p.GossipMessage_MatchedOrder{
				MatchedOrder: &protocols_p2p.MatchedOrderMessageData{
					MessageData:  cr.node.newMessageData(msgData, false),
					MatchedOrder: msg.GetMatchedOrder().MatchedOrder,
				},
			},
		}
	case *protocols_p2p.GossipMessage_TransactionMessage:
		msgData := msg.GetTransactionMessage().MessageData
		sig, err := cr.node.signData(msgData.Data)
		if err != nil {
			cr.node.logger.Errorf("sign error: %v", err)
		}
		msgData.Sign = append(msgData.Sign, sig)
		msg = &protocols_p2p.GossipMessage{
			Message: &protocols_p2p.GossipMessage_TransactionMessage{
				TransactionMessage: &protocols_p2p.TransactionMessageMessageData{
					MessageData:        cr.node.newMessageData(msgData, false),
					TransactionMessage: msg.GetTransactionMessage().TransactionMessage,
				},
			},
		}
	case *protocols_p2p.GossipMessage_OrderMessageData:
		msgData := msg.GetOrderMessageData().MessageData
		sig, err := cr.node.signData(msgData.Data)
		if err != nil {
			cr.node.logger.Errorf("sign error: %v", err)
		}
		msgData.Sign = append(msgData.Sign, sig)
		msg = &protocols_p2p.GossipMessage{
			Message: &protocols_p2p.GossipMessage_OrderMessageData{
				OrderMessageData: &protocols_p2p.OrderMessageData{
					MessageData:       cr.node.newMessageData(msgData, false),
					OrderMessageArray: msg.GetOrderMessageData().OrderMessageArray,
				},
			},
		}
	}
	cr.node.logger.Infof("publishing message")
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	return cr.topic.Publish(cr.ctx, msgBytes)
}

// ReadMessageLoop pulls messages from the pubsub topic and pushes them onto the Messages channel.
func (cr *ChatRoom) ReadMessageLoop() {
	for {
		// defer close(cr.obsRecieveRes)
		select {
		case <-cr.ctx.Done():
			return
		default:
			msg, err := cr.sub.Next(cr.ctx)
			if err != nil {
				close(cr.obsRecieveRes)
				return
			}
			// only forward messages delivered by others
			if msg.ReceivedFrom == cr.self {
				continue
			}
			var cm protocols_p2p.GossipMessage
			err = proto.Unmarshal(msg.Data, &cm)
			if err != nil {
				continue
			}

			// send valid messages onto the Messages channel
			var check bool

			switch cm.Message.(type) {
			case *protocols_p2p.GossipMessage_MatchedOrder:
				check = cr.node.authenticateMessage(&cm, cm.GetMatchedOrder().MessageData)
			case *protocols_p2p.GossipMessage_TransactionMessage:
				check = cr.node.authenticateMessage(&cm, cm.GetTransactionMessage().MessageData)
			case *protocols_p2p.GossipMessage_OrderMessageData:
				check = cr.node.authenticateMessage(&cm, cm.GetOrderMessageData().MessageData)
			default:
				cr.node.logger.Errorf("Wrong message received on p2p")
				check = false
			}
			if check {
				cr.obsRecieveRes <- &cm
			} else {
				continue
			}
		}

	}
}

func (cr *ChatRoom) PublishMessageLoop() {
	for {
		select {
		case input := <-cr.obsSendReq:
			// when order details is passed in inputCh it is published to chatroom after signing and adding message data
			err := cr.Publish(input)
			if err != nil {
				cr.node.logger.Errorf("publish error: %v", err)
			}

		case <-cr.ctx.Done():
			return
		}
	}
}

func topicName(roomName string) string {
	return "chat-room:" + roomName
}
