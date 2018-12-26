package peer

import (
	"encoding/json"
	"log"
	"net"
	"splashp2p/message"
)

type Peer struct {
	Addr      net.UDPAddr
	Behaviour *Behaviour
	Lastmsg   int64
	Conn      net.UDPConn
}

func (p *Peer) Handle(msgRaw []byte) {
	msg := message.Message{}
	err := json.Unmarshal(msgRaw, &msg)
	if err != nil {
		log.Print("peer::"+p.Addr.String(), "[IN] -> ❌ message failed", len(msgRaw), "bytes")
		p.Behaviour.OnMessageFail(EVENT_MESSAGEFAIL, p, err)
	}

	log.Print("peer::"+p.Addr.String(), "[IN] ->", msg.Tag, "->", len(msgRaw), "bytes")

	// Find relevant message handler in peers behaviour
	handler, found := p.Behaviour.MessageHandlers[msg.Tag]
	if !found {
		p.Behaviour.UnknownMessageHandler(msg, p)
		return
	}

	// Handle the message
	handler(msg, p)
	return
}

func (p *Peer) Send(msg message.Message) {
	msgRaw, _ := json.Marshal(&msg)
	length, err := p.Conn.WriteToUDP(msgRaw, &p.Addr)
	if err != nil {
		panic(err)
	}
	log.Print("peer::"+p.Addr.String(), "[OUT] ->", msg.Tag, "->", length, "bytes")
}
