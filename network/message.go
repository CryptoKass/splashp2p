package network

import "encoding/json"

// NetworkMessage is the typical datastruct sent over the network to
// peers; It contains a single tag and payload.
type NetworkMessage struct {
	Tag     string
	Payload []byte
}

func (msg *NetworkMessage) Marshal() ([]byte, error) {
	return json.Marshal(msg)
}
