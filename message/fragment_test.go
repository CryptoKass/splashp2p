package message_test

import (
	"crypto/rand"
	"encoding/hex"
	"testing"
	"time"

	"github.com/CryptoKass/splashp2p/message"
)

func TestDivide(t *testing.T) {
	//Make ultra large payload
	payload := make([]byte, 5000)
	rand.Read(payload)

	//make a new big message
	msg := message.Message{
		Tag:       "test-message",
		Payload:   string(payload),
		Timestamp: time.Now().Unix(),
	}

	//divide message into parts
	parts := message.DivideBigMessage(&msg)
	t.Log("message in", len(parts), "parts")

	if len(parts) != 5 {
		t.Error("message in", len(parts), "parts", "want 5")
	}

	//check each part has correct index and id
	id := hex.EncodeToString(parts[0].ID)
	t.Log("message", hex.EncodeToString(parts[0].ID), "parts:")
	for index, p := range parts {

		if p.Index != int8(index+1) {
			t.Error("incorrect index on part:", p.Index, "want", index)
		}
		if id != hex.EncodeToString(p.ID) {
			t.Error("incorrect id on part:", hex.EncodeToString(p.ID), "want", id)
		}
		t.Logf("  %d/%d:%d bytes", p.Index, p.Max, len(p.Payload))
	}

	res := message.CompileFragments(parts)

	//compare messages
	bufA, _ := res.Marshal()
	bufB, _ := msg.Marshal()

	if string(bufA) != string(bufB) {
		t.Error("Failed to recompile message, had:", len(bufA), "bytes, wanted:", len(bufB), "bytes")
	}

}
