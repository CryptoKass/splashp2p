package message

import (
	"crypto/rand"
	"encoding/hex"
	"hash/fnv"
	"math"
)

var payloadSize = 1000

// MessageFragment - Because of MTU limits some messages that are larger than
// 1472 bytes will fail to arrive properly; In this instance a message can be
// broken up into fragments and ReAssembled at the other end.
//
// Fragments look like a tradition `message.Message` with addtional fields for
// fragment `ID`, `Index`, and `Max` (the number of times the original was
// divided).
type MessageFragment struct {
	Tag       string
	Payload   string
	ID        []byte
	Index     int8
	Max       int8
	Timestamp int64
}

// DivideBigMessage - Divide a message into `MessageFragements`
func DivideBigMessage(msg *Message) (parts []MessageFragment) {
	// generate random id
	id := randID()

	// payload
	payload := []byte(msg.Payload)
	totalSize := len(payload)

	// calculate totalsize
	partCount := int(math.Ceil(float64(totalSize) / float64(payloadSize)))
	index := int8(1)

	// loop payload and append the message parts
	for cursor := 0; cursor < totalSize; {
		partEnd := min(cursor+payloadSize, totalSize)
		section := payload[cursor:partEnd]

		part := MessageFragment{
			Tag:       msg.Tag,
			Payload:   hex.EncodeToString(section),
			ID:        id,
			Index:     index,
			Max:       int8(partCount),
			Timestamp: msg.Timestamp,
		}

		parts = append(parts, part)
		cursor = partEnd
		index++
	}

	return
}

// randID -  generate a random id to be used as `fragment.ID`
func randID() []byte {
	noise := make([]byte, 4)
	rand.Read(noise)

	hasher := fnv.New32a()
	return hasher.Sum(noise)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
