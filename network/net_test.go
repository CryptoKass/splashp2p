package network_test

import (
	"crypto/rand"
	"testing"
	"time"

	"github.com/CryptoKass/splashp2p"
	"github.com/CryptoKass/splashp2p/message"
	"github.com/CryptoKass/splashp2p/network"
	"github.com/CryptoKass/splashp2p/peer"
)

var testVar = 0

func TestNetwork(t *testing.T) {
	//create test set
	testMsgSet := make([]message.Message, 10)

	//populate test set with 10 normal messages
	for i := 0; i < 10; i++ {
		//Make ultra large payload
		payload := make([]byte, 100)
		rand.Read(payload)

		//make a new big message
		msg := message.Message{
			Tag:       "increment-test",
			Payload:   string(payload),
			Timestamp: time.Now().Unix(),
		}

		testMsgSet[i] = msg
	}

	// Create test handlers
	var testhandlers = map[string]peer.MsgHandler{
		"increment-test": func(msg message.Message, p *peer.Peer) {
			testVar++
			t.Logf("recieved msg, testvar=%d/%d", testVar, 10)
		},
	}

	//Create listen behaviour
	behaviour := splashp2p.DefaultBehaviour
	behaviour.MessageHandlers = testhandlers

	running := true

	//Listener
	go func() {
		n := network.CreateNetwork(3333, 1024, behaviour)
		n.Listen()

		for running {
		}
	}()

	//Sender
	go func() {
		n := network.CreateNetwork(3030, 1024, behaviour)
		n.Listen()
		//connect to listener
		n.Connect("127.0.0.1:3333")

		for _, msg := range testMsgSet {
			t.Log("Sending message")
			n.Broadcast(msg)
		}
		for running {
		}
	}()

	//give network time to look over message
	time.Sleep(time.Second * 2)
	//check 10 messages arrived
	if testVar != 10 {
		t.Errorf("some behaviour didnt trigger, testvar was %d, wanted %d", testVar, 10)
	}

	running = false
	return

}
