// Package splashp2p A simple UDP peer-to-peer framework.
package splashp2p

import (
	"log"
	"time"

	"github.com/CryptoKass/splashp2p/message"
	"github.com/CryptoKass/splashp2p/peer"
)

// DefaultBehaviour defines simple ping-pong handlers and logging for
// events.
var DefaultBehaviour = peer.Behaviour{
	OnConnect:             onConnect,
	OnDisconnect:          onDisconnect,
	OnMessageFail:         onMessageFail,
	UnknownMessageHandler: unknownMessageHandler,
	MessageHandlers:       handlers,
}

func pingHandler(msg message.Message, p *peer.Peer) {
	p.Send(message.PongMessage())
}

func pongHandler(msg message.Message, p *peer.Peer) {
	p.Lastmsg = time.Now().Unix()
	log.Print("peer::"+p.Addr.String(), "[STATE] -> Ping-Pong recieved ")
}

func onConnect(e int, p *peer.Peer, err error) {
	log.Print("peer::"+p.Addr.String(), "-> Connected...")
}

func onDisconnect(e int, p *peer.Peer, err error) {
	log.Print("peer::"+p.Addr.String(), "-> Disconnected...")
}

func onMessageFail(e int, p *peer.Peer, err error) {
	log.Print("peer::"+p.Addr.String(), "-> Message failed.")
}

func unknownMessageHandler(msg message.Message, p *peer.Peer) {
	log.Print("peer::"+p.Addr.String(), "->Uknown Message->", msg)
}

var handlers = map[string]peer.MsgHandler{
	"ping": pingHandler,
	"pong": pongHandler,
}
