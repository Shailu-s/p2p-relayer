package p2p

import (
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	cr "github.com/ethereum/go-ethereum/crypto"
	gocrypto "github.com/libp2p/go-libp2p/core/crypto"

	"github.com/gogo/protobuf/proto"
	"github.com/libp2p/go-libp2p/core/peer"
	protocols_p2p "github.com/volmexfinance/relayers/relayer-srv/chat"
	"github.com/volmexfinance/relayers/relayer-srv/utils"
)

// Authenticate incoming p2p message
// message: a protobufs go data object
// data: common p2p message data
func (n *Node) authenticateMessage(message proto.Message, data *protocols_p2p.MessageData) bool {
	//TODO
	n.lock.Lock()
	defer n.lock.Unlock()
	if data.ClientVersion != clientVersion {
		n.logger.Info("Client version mismatched")
		return false
	}
	// store a temp ref to signature and remove it from message data
	// sign is a string to allow easy reset to zero-value (empty string)
	// sign := data.Sign
	// data.Sign = nil

	/*
		// marshall data without the signature to protobufs3 binary format
		// bin, err := proto.Marshal(message)
		// if err != nil {
		// 	n.logger.Println(err, "failed to marshal pb message")
		// 	return false
		// }
	*/

	// restore sig in message data (for possible future use)
	// data.Sign = sign

	// restore peer id binary format from base58 encoded node id data
	peerId, err := peer.Decode(data.NodeId[len(data.NodeId)-1])
	if err != nil {
		n.logger.Println(err, "Failed to decode node id from base58")
		return false
	}

	// verify the data was authored by the signing peer identified by the public key
	// and signature included in the message
	return n.verifyData(data.Data, []byte(data.Sign[len(data.Sign)-1]), peerId, data.NodeAddress[len(data.NodeAddress)-1], data.Chain)
}

// Verify incoming p2p message data integrity
// data: data to verify
// signature: author signature provided in the message payload
// peerId: author peer id from the message payload
// pubKeyData: author public key from the message payload
func (n *Node) verifyData(data []byte, signature []byte, peerIdIn peer.ID, pubKeyData string, chain string) bool {
	//TODO: verify node public key from store
	sig := signature
	// copy(sig, signature[:])
	if sig[64] >= 27 {
		sig[64] -= 27
	}

	pubKeyBytes, err := cr.Ecrecover(data, sig)
	if err != nil {
		n.logger.Warnf("Error authenticating data %v", err)
		return false
	}

	hash0 := crypto.Keccak256Hash(pubKeyBytes[1:])
	checkSummedAddress := strings.ToLower(hash0.String())

	guardianSet := n.GuardianSetMapping[chain]
	for i := 0; i < len(n.GuardianSetMapping[chain]); i++ {
		if checkSummedAddress[26:] == strings.ToLower(guardianSet[i].Address[2:]) {
			if len(guardianSet[i].PeerId) == 0 {
				publicKey, err := gocrypto.UnmarshalSecp256k1PublicKey(pubKeyBytes)
				if err != nil {
					n.logger.Warnf("Error unmashaling pub key data %v", err)
				}

				peerID, err := peer.IDFromPublicKey(publicKey)
				if err != nil {
					n.logger.Warnf("Error generating peer id %v", err)
				}
				if strings.EqualFold(peerID.String(), peerIdIn.String()) {
					guardianSet[i].PeerId = peerID
				} else {
					return false
				}
			}
			if strings.EqualFold(peerIdIn.String(), guardianSet[i].PeerId.String()) {
				return true
			}
			return false
		}
	}
	return false
}

// helper method - generate message data shared between all node's p2p protocols
func (n *Node) newMessageData(msgData *protocols_p2p.MessageData, gossip bool) *protocols_p2p.MessageData {
	// Add protobufs bin data for message author public key
	// this is useful for authenticating  messages forwarded by a node authored by another node

	//TODO: error handling
	privKey, err := utils.GetPrivateKey(n.privKey)
	if err != nil {
		panic("Failed to get public key for sender from local peer store.")
	}
	nodePubKey := cr.FromECDSAPub(&privKey.PublicKey)
	hash0 := crypto.Keccak256Hash(nodePubKey[1:])
	checkSummedAddress := hash0.String()

	return &protocols_p2p.MessageData{ClientVersion: clientVersion,
		NodeId:      append(msgData.NodeId, peer.Encode(n.host.ID())),
		NodeAddress: append(msgData.NodeAddress, checkSummedAddress[26:]),
		Data:        msgData.Data,
		Timestamp:   time.Now().Unix(),
		Gossip:      gossip,
		Sign:        msgData.Sign,
		Chain:       msgData.Chain,
	}
}

// sign binary data using the local node's private key
func (n *Node) signData(data []byte) ([]byte, error) {
	// key := n.host.Peerstore().PrivKey(n.host.ID())
	privKey, _ := utils.GetPrivateKey(n.privKey)
	// sig, err := worker.SignHash(data, privKey)
	sig, err := cr.Sign(data, privKey)
	// res, err := key.Sign(data)
	// println(bytes.Equal(sig, res), "in p2p")
	return sig, err
}
