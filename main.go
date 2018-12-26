package main

import (
	"log"
	"splashp2p/message"
	"splashp2p/network"
	"splashp2p/peer"
	"time"
)

func PingHandler(msg message.Message, p *peer.Peer) {
	p.Send(message.PongMessage())
}

func PongHandler(msg message.Message, p *peer.Peer) {
	p.Lastmsg = time.Now().Unix()
	log.Print("peer::"+p.Addr.String(), "[STATE] -> Ping-Pong recieved ")
}

func OnConnect(e int, p *peer.Peer, err error) {
	log.Print("peer::"+p.Addr.String(), "-> Connected...")
}

func OnMessageFail(e int, p *peer.Peer, err error) {
	log.Print("peer::"+p.Addr.String(), "-> Message failed.")
}

func UnknownMessageHandler(msg message.Message, p *peer.Peer) {
	log.Print("peer::"+p.Addr.String(), "->Uknown Message->", msg)
}

func main() {
	handlers := make(map[string]peer.PeerMsgHandler)
	handlers["ping"] = PingHandler
	handlers["pong"] = PongHandler
	network := network.CreateNetwork(3000, 2048, peer.Behaviour{OnConnect: OnConnect, OnMessageFail: OnMessageFail, MessageHandlers: handlers})
	network.Listen()

	for {

	}
}
