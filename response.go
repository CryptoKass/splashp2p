package splashp2p

// Response contains the outbound message
type Response struct {
	Peer *Peer
	Tag  string
	Body []byte
}

// WriteTag will set the responses tag
func (r *Response) WriteTag(tag string) {
	r.Tag = tag
}

// Write some bytes to the response
func (r *Response) Write(buf []byte) {
	r.Body = append(r.Body, buf...)
}

// Send response as a networkmessage to a peer
func (r *Response) Send() {
	netmsg := NetworkMessage{
		Tag:     r.Tag,
		Payload: r.Body,
	}
	r.Peer.Send(netmsg)
}
