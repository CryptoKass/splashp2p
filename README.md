# splashp2p - SplashLedger:
[![GODOC](https://godoc.org/github.com/CryptoKass/splashp2p?status.svg)](https://godoc.org/github.com/CryptoKass/splashp2p)
 [![Build Status](https://travis-ci.org/CryptoKass/splashp2p.png?branch=master)](https://travis-ci.org/CryptoKass/splashp2p)
[![Coverage Status](https://coveralls.io/repos/github/CryptoKass/splashp2p/badge.svg?branch=master)](https://coveralls.io/github/CryptoKass/splashp2p?branch=master)
*This project is part of the \*offical suite for the Splash Distributed Ledger. This repo is maintained by the Splash Foundation [http://SplashLedger.com](SplashLedger.com)*

<br></br>
# Features:
SplashP2P is a simple peer to peer libary using the built in `net` lib; It was created for the Splash distributed ledger (blockchain).
- UDP
- Inspired by express.js
- ~~Support for big messages upto 255,000 bytes.~~
- No thirdparty depedencies.

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
//create a Network instance
app := splashp2p.NewNetwork()

// define a handler for an inbound request
app.Handle("ping", func(in splashp2p.Request, out splashp2p.Response) {
    if string(in.Body) == "ping" {
        //write new message using the `out` object
        out.WriteTag("pong") //the used to trigger the handler
        out.Write("pong!") //writes to the body
        out.Send() //send the message
    }
})

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


<br></br>
# Contribution:
**If I got something wrong (which I almost certainly have) please let me know:**
- Pull requests welcomed!
- Feedback: cryptokass@gmail.com


<br></br>

---

![OVERVIEW GRAPH](https://i.imgur.com/cUp6QaY.png)


*Readme last updated: 2019.01.14*
