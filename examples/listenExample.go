// Copyright 2018 <kasscrypto@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this
// software and associated documentation files (the "Software"), to deal in the Software
// without restriction, including without limitation the rights to use, copy, modify,
// merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be included in all copies
// or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
// PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF
// CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package examples

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
