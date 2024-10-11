package main

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func newStorage() *Storage {
	opts := StorageOpts{PathTransformFunc: CASPathTransformFunc}
	return NewStorage(opts)
}
func teardown(t *testing.T, s *Storage) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
func TestTransformFunc(t *testing.T) {
	key := "key"
	pathKey := CASPathTransformFunc(key)
	expectedOriginal := "a62f2225bf70bfaccbc7f1ef2a397836717377de"
	assert.Equal(t, "a62f2/225bf/70bfa/ccbc7/f1ef2/a3978/36717/377de", pathKey.PathName)
	assert.Equal(t, expectedOriginal, pathKey.FileName)
}
func TestStorage(t *testing.T) {

	opts := StorageOpts{PathTransformFunc: CASPathTransformFunc}
	s := NewStorage(opts)
	defer teardown(t, s)
	for i := 0; i < 50; i++ {

		key := fmt.Sprintf("key_%d", i)
		data := []byte("data")
		if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		if ok := s.Has(key); !ok {
			t.Error(ok)
		}
		r, err := s.Read(key)
		if err != nil {
			t.Error(err)
		}
		b, _ := io.ReadAll(r)
		fmt.Println(string(b))
		assert.Equal(t, data, b)
		if err := s.Delete(key); err != nil {
			t.Error(err)
		}
		if ok := s.Has(key); ok {
			t.Errorf("Expected to not have key %s", key)
		}
	}
}
func TestDelete(t *testing.T) {
	opts := StorageOpts{PathTransformFunc: CASPathTransformFunc}
	s := NewStorage(opts)
	key := "key"
	data := []byte("data")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
}
