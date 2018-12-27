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

// Package peer Peers are assigned those splashp2p will communicate with.
package peer

import (
	"encoding/json"
	"log"
	"net"
	"time"

	"github.com/CryptoKass/splashp2p/message"
)

// Peer - Peer manages commincation to and from a single peer.
// Inbound messages are handled as json `[]byte` and in the
// structure found in message.Message.
//
// Generally Peer objects are created by `network.Net` when a
// message is recieved from a new IP.
//
// TODO: add support for a custom message interface.
// TODO: add custom state struct.
// TODO: add pointer to the network.Net parent.
type Peer struct {
	Addr      net.UDPAddr
	Behaviour *Behaviour
	Lastmsg   int64
	Conn      net.UDPConn

	//Support for big messages:
	bigDecay     time.Duration
	FragmentsIn  map[string]FragmentIn
	FragmentsOut map[string]FragmentOut
}

// Handle - Raw inbound message bytes are passed into `handle`.
// The bytes are unmarshaled and the messsage.Message is handled
// using p.Behaviour.MessageHandlers[msg.Tag].
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

// Send - will write a message.Message to this peers UDP
// address. If the message is bigger than 1400 bytes
// SendBigMessage will be used
func (p *Peer) Send(msg message.Message) {
	msgRaw, _ := json.Marshal(&msg)

	//check if message is too big
	if len(msgRaw) >= 1400 {
		p.SendBigMessage(msg)
		return
	}

	length, err := p.Conn.WriteToUDP(msgRaw, &p.Addr)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("peer::"+p.Addr.String(), "[OUT] ->", msg.Tag, "->", length, "bytes")
}

// sendFragment - will write a `message.MessageFragment`
// to this peers UDP address. You should use SendBigMessage()
func (p *Peer) sendFragment(msg message.MessageFragment) {
	msgRaw, _ := json.Marshal(&msg)
	length, err := p.Conn.WriteToUDP(msgRaw, &p.Addr)
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("peer::"+p.Addr.String(), "[OUT] ->", msg.Tag, msg.Index, "/", msg.Max, "->", length, "bytes")
}

// SendBigMessage - will divide a large message into
// `message.MessageFragment` and will send them indivually
// to the peers UDP address.
//
// SendBigMessage is designed for messages larger than
// 1400 bytes but smaller than 255,000 bytes
func (p *Peer) SendBigMessage(msg message.Message) {
	// this may take some time, best to do it asynchronusly
	go func() {
		// Divide messages
		parts := message.DivideBigMessage(&msg)

		// make outgoing fragment manager incase a message gets
		// lost and the peer can request it
		fragOut := FragmentOut{
			id:     string(parts[0].ID),
			parent: p,
			parts:  parts,
		}
		fragOut.timeout = time.AfterFunc(p.bigDecay*2, fragOut.Decay)
		p.FragmentsOut[fragOut.id] = fragOut

		// BigFragment
		for _, v := range parts {
			p.sendFragment(v)
			time.Sleep(50 * time.Millisecond)
		}

	}()
}
