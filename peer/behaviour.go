package peer

import (
	"splashp2p/message"
)

type Behaviour struct {
	OnConnect             PeerEvent
	OnDisconnect          PeerEvent
	OnMessageFail         PeerEvent
	UnknownMessageHandler PeerMsgHandler
	MessageHandlers       map[string]PeerMsgHandler
}

type PeerEvent func(e int, p *Peer, err error)

type PeerMsgHandler func(msg message.Message, p *Peer)
