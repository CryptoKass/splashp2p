# splashp2p - SplashLedger:
[![GoDoc] https://godoc.org/github.com/CryptoKass/splashp2p?status.svg](https://godoc.org/github.com/CryptoKass/splashp2p)

*This project is part of the \*offical suite for the Splash Distributed Ledger. This repo is maintained by the Splash Foundation [http://SplashLedger.com](SplashLedger.com)*

#Getting stated:
Download:
```shell
go get splashp2p
```
If you do not have the go command on your system, you need to [Install Go](http://golang.org/doc/install) first


# Features:
SplashP2P is a simple peer to peer libary using the built in `net` lib; It was created for the Splash distributed ledger (blockchain).
- UDP
- No thirdparty depedencies
- Configurable Behaviour.

# Examples:
Create your peer to peer network with `network.CreateNetwork(port, message, behavour)`.
Here is an example using the DefaultBehaviour which is a ping-pong example:
```golang
func main() {
	p2pnetwork := network.CreateNetwork(3000/*port*/, 2048/*message size*/, splashp2p.DefaultBehaviour/*behavour*/)
	p2pnetwork.Listen()
	for {//forever}
}
```

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

For more examples please see the examples subdirectory or look at the godoc.

# Todo:
There are a range of things I have planned for this lib:
- Customizable peer state
- Implement peer.Disconnect (will require peer to have a pointer to *Net parent)
- Add peer to timers -> timeout ect.
- Improve documentation

# Documentaion:
Visit [GoDoc](https://godoc.org/github.com/CryptoKass/splashp2p)
Package splashp2p A simple UDP peer-to-peer framework.



# Contribution:
**If I got something wrong (which I almost certainly have) please let me know:**
- Pull requests welcomed!
- Feedback: cryptokass@gmail.com

*Readme last updated: 2018.12.26*
