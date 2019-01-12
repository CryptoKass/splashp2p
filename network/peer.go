package network

import (
	"encoding/json"
	"log"
	"net"
)

type Peer struct {
	Addr    net.UDPAddr
	Lastmsg int64
	Conn    net.UDPConn
}

// Send - will write a message.Message to this peers UDP
// address. If the message is bigger than 1400 bytes
// SendBigMessage will be used
func (p *Peer) Send(netMsg NetworkMessage) {
	msgRaw, _ := json.Marshal(&netMsg)

	//check if message is too big
	if len(msgRaw) >= 1400 {
		//p.SendBigMessage(msg)
		return
	}
	_, err := p.Conn.WriteToUDP(msgRaw, &p.Addr)
	if err != nil {
		log.Println(err)
		return
	}

}
