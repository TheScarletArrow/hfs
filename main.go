package main

import (
	"gitlab.com/TheScarletArrow/hfs/p2p"
	"log"
)

func makeServer(listenAddr string, nodes ...string) *Server {
	tcpTrabsportOpts := p2p.TCPTransportOps{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	transport := p2p.NewTCPTransport(tcpTrabsportOpts)

	serverOpts := ServerOpts{
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         transport,
		BootstrapNodes:    nodes,
	}
	s := NewServer(serverOpts)
	tcpTrabsportOpts.OnPeer = s.OnPeer
	return s

}

func main() {
	s1 := makeServer(":3000", "")
	s2 := makeServer(":4000", ":3000")
	go func() {
		log.Fatal(s1.Start())
	}()

	s2.Start()
}
