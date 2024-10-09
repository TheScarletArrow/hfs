package main

import (
	"github.com/stretchr/testify/assert"
	"gitlab.com/TheScarletArrow/hfs/p2p"
	"testing"
)

func TestTCPTransport(t *testing.T) {
	opts := p2p.TCPTransportOps{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(opts)

	assert.Equal(t, tr.ListenAddr, opts.ListenAddr)

	assert.Nil(t, tr.ListenAndAccept())

}
