# splashp2p - SplashLedger:
[![GODOC](https://godoc.org/github.com/CryptoKass/splashp2p?status.svg)](https://godoc.org/github.com/CryptoKass/splashp2p)
 [![Build Status](https://travis-ci.org/CryptoKass/splashp2p.png?branch=master)](https://travis-ci.org/CryptoKass/splashp2p)
*This project is part of the \*offical suite for the Splash Distributed Ledger. This repo is maintained by the Splash Foundation [http://SplashLedger.com](SplashLedger.com)*

<br></br>
# Features:
SplashP2P is a simple peer to peer libary using the built in `net` lib; It was created for the Splash distributed ledger (blockchain).
- UDP
- Support for big messages upto 255,000 bytes.
- No thirdparty depedencies.
- Configurable Behaviour.

<br></br>
# Getting stated:
Download:
```shell
go get github.com/CryptoKass/splashp2p
```
If you do not have the go command on your system, you need to [Install Go](http://golang.org/doc/install) first

<br></br>
# Usage:
```golang
func main() {
    // load default logic
    logic := splashp2p.DefaultBehaviour
        
    // add hello world message handler
   logic.MessageHandlers["hello"] = helloworldHandler
    
    // create p2p manager object
    p2p := network.CreateNetwork(3000/*port*/, 1024/*message size*/, logic/*behavour*/)

    //begin listening over udp
    p2p.Listen()

    for {//forever}
}

func helloworldHandler(msg message.Message, p *peer.Peer) {
    out := message.Message{
        Tag:       "hello-world-response",
        Payload:   "Hello World!",
        Timestamp: time.Now().Unix(),
    }
     p.Send(out)
}
```

<br></br>
# Examples:
Below are some examples you may find useful:

<br></br>
### Listen to network: *Net.CreateNetwork 
Create your peer to peer network with `network.CreateNetwork(port, message, behavour)`.
Here is an example using the DefaultBehaviour which is a ping-pong example:
```golang
func main() {
	p2pnetwork := network.CreateNetwork(3000/*port*/, 1024/*message size*/, splashp2p.DefaultBehaviour/*behavour*/)
	p2pnetwork.Listen()
	for {//forever}
}
```

<br></br>
### Connect to a peer
Connect to a givens peers IP. Assuming the peer is listening, with `network.Connect(ip)`.
```golang
p2pnetwork.Listen() //must be listening before you can connect to a peer
p2pnetwork.Connect("some.peer:6677")
```

<br></br>
### Custom logic: *Behaviour.MessageHandlers
You can define how you will handle messages and events by creating your own `peer.Behaviour`.
(Example:)
```golang
// message handler must conform to `func(msg message.Message, p *Peer)`
func helloworldHandler(msg message.Message, p *peer.Peer) {
    out := message.Message{
        Tag:       "hello-world-response",
        Payload:   "Hello World!",
        Timestamp: time.Now().Unix(),
    }
    p.Send(out)
}

customBehaviour := peer.Behaviour{...}
customBehaviour.MessageHandlers["hello"] = helloworldHandler

```

<br></br>
### Broadcast: *Net.Broadcast
You can broadcast a message to all connected peers using `*Net.Broadcast`
```golang
// create a network.Net to interface with peers
p2pnetwork := p2pnetwork := network.CreateNetwork(...)
//connect to some peers
p2pnetwork.Connect("somepeer:6677")
// define a message to broadcast
msg := message.Message{Tag:"hello-world", Timestamp: time.Now().Unix()}
// broadcast the message
p2pnetwork.Broadcast(msg)
```

For more examples please see the examples subdirectory or look at the godoc.

<br></br>
# Todo:
There are a range of things I have planned for this lib:
- Add test coverage.
- Customizable peer state
- Implement peer.Disconnect (will require peer to have a pointer to *Net parent)
- Add peer to timers -> timeout ect.
- Improve documentation

<br></br>
# Documentaion:
Visit [GoDoc](https://godoc.org/github.com/CryptoKass/splashp2p):
Package splashp2p A simple UDP peer-to-peer framework.



<br></br>
# Contribution:
**If I got something wrong (which I almost certainly have) please let me know:**
- Pull requests welcomed!
- Feedback: cryptokass@gmail.com


<br></br>

---

![OVERVIEW GRAPH](https://i.imgur.com/cUp6QaY.png)


*Readme last updated: 2018.12.27*
