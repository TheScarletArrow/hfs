package main

import (
	"fmt"
	"gitlab.com/TheScarletArrow/hfs/p2p"
	"log"
)

func OnPeer(peer p2p.Peer) error {
	fmt.Println("doint some")
	return nil
}
func main() {
	ops := p2p.TCPTransportOps{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}
	tr := p2p.NewTCPTransport(ops)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()

	if err := (tr.ListenAndAccept()); err != nil {
		log.Fatal(err)
	}
	select {}
}
