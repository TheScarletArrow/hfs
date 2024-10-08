package p2p

import "net"

// Message represent any arbitrary data is being sent over the
// each transport between two nodes
type Message struct {
	From    net.Addr
	Payload []byte
}
