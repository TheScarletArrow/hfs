package main

import (
	"gitlab.com/TheScarletArrow/hfs/p2p"
	"log"
)

func main() {
	ops := p2p.TCPTransportOps{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(ops)

	if err := (tr.ListenAndAccept()); err != nil {
		log.Fatal(err)
	}
	select {}
}
