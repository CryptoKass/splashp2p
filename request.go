package splashp2p

// Request contains a inbound message, and
// associated data.
type Request struct {
	Sender   *Peer
	Tag      string
	Body     []byte
	Timstamp int64
	Vars     map[string]interface{}
}
