package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestTransformFunc(t *testing.T) {
	key := "key"
	pathKey := CASPathTransformFunc(key)
	expectedOriginal := "a62f2225bf70bfaccbc7f1ef2a397836717377de"
	assert.Equal(t, "a62f2/225bf/70bfa/ccbc7/f1ef2/a3978/36717/377de", pathKey.PathName)
	assert.Equal(t, expectedOriginal, pathKey.FileName)
}
func TestStorage(t *testing.T) {

	opts := StorageOpts{PathTransformFunc: CASPathTransformFunc}
	s := NewStore(opts)

	key := "key"
	data := []byte("data")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}
	r, err := s.read(key)
	if err != nil {
		t.Error(err)
	}
	b, _ := io.ReadAll(r)
	fmt.Println(string(b))
	assert.Equal(t, data, b)
	s.delete(key)
}
func TestDelete(t *testing.T) {
	opts := StorageOpts{PathTransformFunc: CASPathTransformFunc}
	s := NewStore(opts)
	key := "key"
	data := []byte("data")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.delete(key); err != nil {
		t.Error(err)
	}
}
