package p2p

import "net"

// RPC represent any arbitrary data is being sent over the
// each transport between two nodes
type RPC struct {
	From    net.Addr
	Payload []byte
}
