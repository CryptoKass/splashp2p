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
	"splashp2p/message"
)

// Behaviour Describes how a peer will response to events
// and handle inbound messages.
type Behaviour struct {
	// Triggered when Connecting to a peer.
	// This code will run wether the peer
	// recieves the message or not.
	OnConnect Event
	// This will trigger just before peer.Disconnect()
	// is run.
	OnDisconnect Event
	// This will trigger when an inbound message fails
	// to be parsed.
	OnMessageFail Event
	// Describes how to handle a message that is not in
	// MessageHandlers
	UnknownMessageHandler MsgHandler
	// Handle messages - the message.Tag is the
	// key and MsgHandler is the value.
	MessageHandlers map[string]MsgHandler
}

// Event - `Peer.Event` is called by `peer.Peer` to handle events
// such as OnConnect, OnDisconnect and OnMessageFail.
type Event func(e int, p *Peer, err error)

// MsgHandler - `MessageHandler` is called by `peer.Peer` to
// handle inbound message.
type MsgHandler func(msg message.Message, p *Peer)
