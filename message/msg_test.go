package message_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/CryptoKass/splashp2p/message"
)

func TestEncoding(t *testing.T) {
	//test data
	testTag := "test-tag"
	testPayload := "testing 1 2 3 4 5 6 7 8"
	testTime := time.Now().Unix()

	//create a message
	msg := message.Message{
		Tag:       testTag,
		Payload:   testPayload,
		Timestamp: testTime,
	}

	//encode message
	enc, err := msg.Marshal()
	if err != nil {
		t.Error("failed to encoded msg, err dump ->", err)
	}

	//decode message
	dec := message.Message{}
	err = json.Unmarshal(enc, &dec)
	if err != nil {
		t.Error("failed to dencode msg, err dump ->", err)
	}

	//compare messages
	if dec.Tag != testTag {
		t.Errorf("incorrect decoded Tag, have %s, want %s", dec.Tag, testTag)
	}
	if dec.Payload != testPayload {
		t.Errorf("incorrect decoded Tag, have %s, want %s", dec.Payload, testPayload)
	}
	if dec.Timestamp != testTime {
		t.Errorf("incorrect decoded Tag, have %d, want %d", dec.Timestamp, testTime)
	}

	t.Log("Decoded Message:")
	t.Log(" ", dec)
}
