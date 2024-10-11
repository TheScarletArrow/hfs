package main

import (
	"fmt"
	"gitlab.com/TheScarletArrow/hfs/p2p"
	"log"
	"sync"
)

type ServerOpts struct {
	ListenAddr        string
	StorageRoot       string
	PathTransformFunc PathTransformFunc
	Transport         p2p.Transport
	TransportOpts     p2p.TCPTransportOps
	BootstrapNodes    []string
}

type Server struct {
	ServerOpts

	peerLock sync.Mutex
	peers    map[string]p2p.Peer

	storage *Storage
	quitch  chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	storageOpts := StorageOpts{
		Root:              opts.StorageRoot,
		PathTransformFunc: opts.PathTransformFunc}
	return &Server{
		storage:    NewStorage(storageOpts),
		ServerOpts: opts,
		quitch:     make(chan struct{}),
		peers:      make(map[string]p2p.Peer),
	}
}

func (s *Server) Stop() {
	close(s.quitch)
}

func (s *Server) OnPeer(peer p2p.Peer) error {
	s.peerLock.Lock()
	defer s.peerLock.Unlock()
	s.peers[peer.RemoteAddr().String()] = peer
	log.Printf("connected to peer: %s", peer)
	return nil
}

func (s *Server) loop() {
	defer func() {
		log.Println("server stopped due to user action")
		s.Transport.Close()
	}()
	for {
		select {
		case msg := <-s.Transport.Consume():
			fmt.Println(msg)
		case <-s.quitch:
			return
		}
	}
}

func (s *Server) bootstrapNetwork() error {
	for _, addr := range s.BootstrapNodes {
		if len(addr) == 0 {
			continue
		}
		go func(addr string) {
			log.Printf("bootstrapping network: %s", addr)
			if err := s.Transport.Dial(addr); err != nil {
				log.Printf("failed to bootstrap network: %s", err)
			}
		}(addr)
	}

	return nil
}

func (s *Server) Start() error {
	if err := s.Transport.ListenAndAccept(); err != nil {
		return err
	}

	s.bootstrapNetwork()

	s.loop()
	return nil
}
