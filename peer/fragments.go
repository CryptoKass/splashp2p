package peer

import (
	"time"

	"github.com/CryptoKass/splashp2p/message"
)

// FragmentIn - An 'inbox' to store parts of an inbound big message.
// output should decay on a given timer.
type FragmentIn struct {
	id      string
	parent  *Peer
	parts   []message.MessageFragment
	max     int8
	inv     []int8
	timeout *time.Timer
}

// Add - Append a `new MessageFragment` to this inbox.
// Successfully adding an unknown message will reset
// the timer.
// If the inbox is complete. Compile and decay.
func (inbox *FragmentIn) Add(msg message.MessageFragment) {
	// check if in inv
	for _, v := range inbox.inv {
		if v == msg.Index {
			return // we already have this part
		}
	}
	// check is smaller than max
	if msg.Index > inbox.max {
		return
	}
	// register
	inbox.parts = append(inbox.parts, msg)
	inbox.inv = append(inbox.inv, msg.Index)

	//update timer
	if !inbox.timeout.Stop() {
		<-inbox.timeout.C //drain timeout channel
	}
	inbox.timeout = time.AfterFunc(inbox.parent.bigDecay, inbox.Decay)

	// is inv compltete -> compile and decay
	if len(inbox.inv) < int(inbox.max) {
		return
	}

}

// Compile - Merge the payload of the `inbox.parts` and
// handle the resulting `message.Message`.
func (inbox *FragmentIn) Compile() {

	msg := message.CompileFragments(inbox.parts)
	buf, _ := msg.Marshal()

	//handle
	inbox.parent.Handle(buf)

	//decay
	inbox.Decay()
}

// Decay - Destroy self by removing it from the parents
// map.
func (inbox *FragmentIn) Decay() {
	delete(inbox.parent.FragmentsIn, inbox.id)

	//destroy timer
	if !inbox.timeout.Stop() {
		<-inbox.timeout.C //drain timeout channel
	}
}

// FragmentOut - An 'outbox' to store parts of an outbound big message.
// Outbox will decay within a given timer.
type FragmentOut struct {
	id      string
	parent  *Peer
	parts   []message.MessageFragment
	timeout *time.Timer
}

// Decay - Destroy self by removing it from the parents
// map.
func (outbox *FragmentOut) Decay() {
	delete(outbox.parent.FragmentsOut, outbox.id)

	//destroy timer
	if !outbox.timeout.Stop() {
		<-outbox.timeout.C //drain timeout channel
	}
}

// Resend - Retrive and 'resend' a message part to the peer.
func (outbox *FragmentOut) Resend(index int, resetTimer bool) {
	if index >= len(outbox.parts) {
		return
	}

	if resetTimer {
		outbox.timeout = time.AfterFunc(outbox.parent.bigDecay, outbox.Decay)
		if !outbox.timeout.Stop() {
			<-outbox.timeout.C //drain timeout channel
		}
	}

	part := outbox.parts[index]
	outbox.parent.sendFragment(part)
}
