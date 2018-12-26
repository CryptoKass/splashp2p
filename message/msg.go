package message

import (
	"time"
)

type Message struct {
	Tag       string
	Payload   string
	Timestamp int64
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
