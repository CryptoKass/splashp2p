package network

import (
	"fmt"
	"net"
	"splashp2p/message"
	"splashp2p/peer"
	"strconv"
	"time"
)

// Net is a structure used to manage peer communication including
// incoming and outgoing messages. Net uses UDP to communicates.
// For each peer that connects a network.Peer is associated with
// them and is the conext used to Handle message to and from that
// peer.
//
// Net should be constructd with:
// 	net := CreateNetwork()
//
// You can begin listening with
// 	net.Listen()
//
// You can automattically connect to peers using:
// 	net.Connect("splashnode.io:6677")
//
type Net struct {
	// ConnectedPeer - A list of peers with active connnections.
	ConnectedPeers map[string]peer.Peer
	// Address - The address of the local node.
	Address net.UDPAddr
	// Conn - The UDP 'connection' used to send messages.
	Conn *net.UDPConn
	// Port - The port that this node is listening on
	// by default this is 6677.
	Port int
	// MaxMessageSize - Max size of message that this node is willing
	// To recieve.
	MaxMessageSize int
	// Listening - can be set to false to stop the
	// LIsten loop, if it is running.
	Listening bool
	// Behaviour: Default peer behaviour
	Behaviour peer.Behaviour
}

// CreateNetwork is used to initialize a network.Net object.
// see network.Net, network.*Net.Listen, network.*Net.Connect
func CreateNetwork(port int, maxMessageSize int, peerBehaviour peer.Behaviour) Net {
	return Net{
		ConnectedPeers: make(map[string]peer.Peer),
		Address:        net.UDPAddr{Port: port},
		Port:           port,
		MaxMessageSize: maxMessageSize,
		Behaviour:      peerBehaviour,
	}
}

// Listen - Listen will await for incomming UDP traffic to the
// local address at the port specifcied by n.Port.
//
// Inbound messages are passed to peer.Handle().
//
// If input is from an unknown peer, a new network.Peer object
// is created and associated with its IP.
func (n *Net) Listen() {
	var err error
	laddr := &n.Address

	//Listen
	n.Conn, err = net.ListenUDP("udp4", laddr)
	if err != nil {
		panic(err)
	}

	fmt.Println("[LISTENING]", "->", laddr)

	n.Listening = true
	go func() {
		for n.Listening {
			buf := make([]byte, n.MaxMessageSize)
			length, raddr, err := n.Conn.ReadFromUDP(buf)
			if err != nil {
				panic(err)
			}

			//Get peer
			cpeer, known := n.ConnectedPeers[raddr.String()]

			//if peer is unknown createa new peer:
			if !known {
				cpeer = peer.Peer{
					Addr:      *raddr,
					Lastmsg:   time.Now().Unix(),
					Conn:      *n.Conn,
					Behaviour: &n.Behaviour,
				}
				n.ConnectedPeers[raddr.String()] = cpeer
				//cpeer.Behaviour.OnConnect(peer.EVENT_CONNECT, &cpeer, nil)
			}

			//pass message to peer object to handle
			cpeer.Handle(buf[:length])
		}
	}()
}

// Connect can be used to connect to a previously unknown peer.
// If a peer is already in n.ConnectedPeers then a ping message
// is sent and `false` is returned.
func (n *Net) Connect(address string) bool {

	// parse address string into net.UDPAddr
	ip, portString, _ := net.SplitHostPort(address)
	port, _ := strconv.ParseInt(portString, 10, 64)
	addr := net.UDPAddr{IP: net.ParseIP(ip), Port: int(port)}

	// check if peer is already connected
	if cpeer, known := n.ConnectedPeers[address]; known {
		fmt.Println("[CONNECT_ERROR] peer is already known.")
		cpeer.Send(message.PingMessage())
		return false
	}

	//Create new peer object
	newpeer := peer.Peer{
		Addr:      addr,
		Lastmsg:   time.Now().Unix(),
		Conn:      *n.Conn,
		Behaviour: &n.Behaviour,
	}

	//Ping the peer
	n.ConnectedPeers[address] = newpeer
	newpeer.Send(message.PingMessage())

	// !DONE
	return true
}