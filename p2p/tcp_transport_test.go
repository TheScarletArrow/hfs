package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := "127.0.0.1:0"
	tr := NewTCPTransport(listenAddr)

	assert.Equal(t, tr.listenAddress, listenAddr)

	assert.Nil(t, tr.ListenAndAccept())

	select {}
}
