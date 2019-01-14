package splashp2p

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var Logger = log.New(os.Stdout, "splashp2p >> ", 1)
var Debug = false
var LogOutput = os.Stdout

type Middleware func(in *Request, out *Response) bool
type Handler func(in Request, out Response)

type Network struct {
	Address        net.UDPAddr
	Conn           *net.UDPConn
	Peers          map[string]*Peer
	Listening      bool
	handlers       map[string]func(in Request, out Response)
	middleware     []Middleware
	MaxMessageSize int
}

func NewNetwork() *Network {
	return &Network{
		Peers:          make(map[string]*Peer),
		handlers:       make(map[string]func(in Request, out Response)),
		middleware:     make([]Middleware, 0),
		MaxMessageSize: 1400,
	}
}

// Listen - Listen will await for incomming UDP traffic to the
// local address at the port specifcied by n.Port.
// Inbound messages are passed to peer.Handle();
// If input is from an unknown peer, a new network.Peer object
// is created and associated with its IP.
func (n *Network) Listen(port int) {

	if !Debug {
		Logger.SetFlags(0)
		Logger.SetOutput(ioutil.Discard)
	} else {
		Logger.SetOutput(LogOutput)
	}

	n.Address = net.UDPAddr{Port: port}

	var err error
	laddr := &n.Address

	//Listen
	n.Conn, err = net.ListenUDP("udp4", laddr)
	if err != nil {
		panic(err)
	}

	Logger.Print("listening on ", laddr)

	n.Listening = true
	go func() {
		for n.Listening {
			// Read inbound message
			buf := make([]byte, n.MaxMessageSize)
			length, raddr, err := n.Conn.ReadFromUDP(buf)
			if err != nil {
				if n.Conn == nil {
					n.Listening = false
				}
				Logger.Print("udpreaderror - failed to read from udp, dumping error below")
				Logger.Print(err)
				continue
			}

			// Get peer object
			cpeer, known := n.Peers[raddr.String()]

			// if peer is unknown createa new peer:
			if !known {
				cpeer = &Peer{
					Addr:    *raddr,
					Lastmsg: time.Now().Unix(),
					Conn:    *n.Conn,
				}
				n.Peers[raddr.String()] = cpeer
				//cpeer.Behaviour.OnConnect(peer.EVENT_CONNECT, &cpeer, nil)
			}

			//pass message to peer object to handle
			n.formatAndHandleRequest(cpeer, buf[:length])
		}
	}()
}

func (n *Network) formatAndHandleRequest(peer *Peer, buf []byte) {
	var netMsg NetworkMessage
	err := json.Unmarshal(buf, &netMsg)
	if err != nil {
		Logger.Println("err - bad request, bytes failed to parse", err)
		return
	}

	//format request
	in := Request{
		Sender: peer,
		Tag:    netMsg.Tag,
		Body:   netMsg.Payload,
		Vars:   make(map[string]interface{}),
	}
	//format response
	out := Response{
		Peer: peer,
		Tag:  "",
		Body: []byte{},
	}

	//pass through middleware
	for i := 0; i < len(n.middleware); i++ {
		if !n.middleware[i](&in, &out) {
			return
		}
	}

	//handle
	if handler, known := n.handlers[in.Tag]; known {
		handler(in, out)
		return
	}

	Logger.Printf("unknown request '%s' recieved from peer %s.", in.Tag, in.Sender.Addr.String())
	return
}

// Handle registers a new handler for a given command. The handler
// is a function in the form `func(in Request, out Response)`
func (n *Network) Handle(command string, handler Handler) {
	n.handlers[command] = handler
}

// AddMiddleware will add a new Middleware function to the stack.
// Middleware will be called on a valid request, before a message is passed to the handler;
// If a middleware func doesnt return true, then request will be discarded.
//
// Middleware is a function in the form `func(in Request, out Response) bool`
func (n *Network) AddMiddleware(middleware Middleware) {
	n.middleware = append(n.middleware, middleware)
}

// Connect can be used to connect to a previously unknown peer.
// If a peer is already in n.Peers then a ping message
// is sent and `false` is returned.
func (n *Network) Connect(address string) bool {

	// parse address string into net.UDPAddr
	ip, portString, _ := net.SplitHostPort(address)
	port, _ := strconv.ParseInt(portString, 10, 64)
	addr := net.UDPAddr{IP: net.ParseIP(ip), Port: int(port)}

	// check if peer is already connected
	if cpeer, known := n.Peers[address]; known {
		Logger.Print("connectionerror - peer is already known.", cpeer.Addr.String())
		return false
	}

	//Create new peer object
	newpeer := Peer{
		Addr:    addr,
		Lastmsg: time.Now().Unix(),
		Conn:    *n.Conn,
	}

	//Ping the peer
	n.Peers[address] = &newpeer
	//newpeer.Send(message.PingMessage())

	// !DONE
	return true
}

// Broadcast - broadcast a message to all peers
func (n *Network) Broadcast(msg NetworkMessage) {

	// broadcast asynchronusly because `net.WriteToUDP` is slow...
	go func() {

		// Call .Send(msg) on all connected peers
		for _, cpeer := range n.Peers {
			cpeer.Send(msg)
		}

		// done!
	}()

}
