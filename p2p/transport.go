package p2p

// Peer is interface that represents remote node
type Peer interface {
}

// Transport is anything that handles communication
// between nodes in network. This can be of a form (TCP UDP WebSockets)
type Transport interface {
	ListenAndAccept() error
}
