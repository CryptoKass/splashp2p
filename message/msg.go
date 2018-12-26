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

// Package message Contains the message structure that is used by splashp2p
package message

import (
	"encoding/json"
	"time"
)

type Message struct {
	Tag       string
	Payload   string
	Timestamp int64
}

// Marshal - convert a message to json ([]byte)
func (m *Message) Marshal() ([]byte, error) {
	return json.Marshal(&m)
}

// PingMessage will generate a generic ping message
// Ping messages are often used to measure latency
// or check that a connection is still alive.
// Peers who send a ping message will often await
// a pong response message.
func PingMessage() Message {
	return Message{
		Tag:       "ping",
		Payload:   "ping",
		Timestamp: time.Now().Unix(),
	}
}

// PongMessage will generate a generic pong message
// Pong messages are often used to reply to ping
// messages.
// Peers will often reply to a inbound ping message
// with a pong message.
func PongMessage() Message {
	return Message{
		Tag:       "pong",
		Payload:   "pong",
		Timestamp: time.Now().Unix(),
	}
}
