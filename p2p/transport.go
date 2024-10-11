package p2p

import "net"

// Peer is interface that represents remote node

// Transport is anything that handles communication
// between nodes in network. This can be of a form (TCP UDP WebSockets)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
	Dial(string) error
}

type Peer interface {
	Send([]byte) error
	RemoteAddr() net.Addr
	Close() error
}
