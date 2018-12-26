# github.com/CryptoKass/splashp2p - SplashLedger:
[![GODOC](https://godoc.org/github.com/CryptoKass/github.com/CryptoKass/splashp2p?status.svg)](https://godoc.org/github.com/CryptoKass/github.com/CryptoKass/splashp2p)
 [![Build Status](https://travis-ci.org/CryptoKass/splashp2p.png?branch=master)](https://travis-ci.org/CryptoKass/splashp2p)
*This project is part of the \*offical suite for the Splash Distributed Ledger. This repo is maintained by the Splash Foundation [http://SplashLedger.com](SplashLedger.com)*

# Getting stated:
Download:
```shell
go get github.com/CryptoKass/splashp2p
```
If you do not have the go command on your system, you need to [Install Go](http://golang.org/doc/install) first


# Features:
SplashP2P is a simple peer to peer libary using the built in `net` lib; It was created for the Splash distributed ledger (blockchain).
- UDP
- No thirdparty depedencies
- Configurable Behaviour.


# Examples:
Below are some examples you may find useful:


### Listen to network: *Net.CreateNetwork 
Create your peer to peer network with `network.CreateNetwork(port, message, behavour)`.
Here is an example using the DefaultBehaviour which is a ping-pong example:
```golang
func main() {
	p2pnetwork := network.CreateNetwork(3000/*port*/, 2048/*message size*/, splashp2p.DefaultBehaviour/*behavour*/)
	p2pnetwork.Listen()
	for {//forever}
}
```


### Connect to a peer
Connect to a givens peers IP. Assuming the peer is listening, with `network.Connect(ip)`.
```golang
p2pnetwork.Listen() //must be listening before you can connect to a peer
p2pnetwork.Connect("some.peer:6677")
```


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


# Todo:
There are a range of things I have planned for this lib:
- Add test coverage.
- Customizable peer state
- Implement peer.Disconnect (will require peer to have a pointer to *Net parent)
- Add peer to timers -> timeout ect.
- Improve documentation


# Documentaion:
Visit [GoDoc](https://godoc.org/github.com/CryptoKass/github.com/CryptoKass/splashp2p)
Package splashp2p A simple UDP peer-to-peer framework.




# Contribution:
**If I got something wrong (which I almost certainly have) please let me know:**
- Pull requests welcomed!
- Feedback: cryptokass@gmail.com

</br>
*Readme last updated: 2018.12.26*
