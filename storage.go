package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
)

type PathTransformFunc func(string) PathKey

type PathKey struct {
	PathName string
	FileName string
}

func (p PathKey) FirstPathName() string {
	paths := strings.Split(p.PathName, "/")
	if len(paths) == 0 {
		return ""
	}
	return paths[0]
}

var DefaultPathTransformFunc = func(key string) string {
	return key
}

func CASPathTransformFunc(key string) PathKey {
	hash := sha1.Sum([]byte(key))
	hashString := hex.EncodeToString(hash[:])

	blockSize := 5
	sliceLen := len(hashString) / blockSize

	paths := make([]string, sliceLen)

	for i := 0; i < sliceLen; i++ {
		from, to := i*blockSize, (i+1)*blockSize
		paths[i] = hashString[from:to]
	}
	return PathKey{
		PathName: strings.Join(paths, "/"),
		FileName: hashString,
	}
}

type StorageOpts struct {
	PathTransformFunc PathTransformFunc
}

type Storage struct {
	StorageOpts
}

func NewStore(opts StorageOpts) *Storage {
	return &Storage{StorageOpts: opts}
}
func (s *Storage) has(key string) bool {
	pathKey := s.PathTransformFunc(key)

	_, err := os.Stat(pathKey.fullPath())

	if !errors.Is(err, fs.ErrNotExist) {
		return false
	}
	return true

}
func (s *Storage) delete(key string) error {

	pathKey := s.PathTransformFunc(key)
	defer func() {
		log.Printf("Deleted %s from disk", pathKey.FileName)
	}()
	if err := os.RemoveAll(pathKey.fullPath()); err != nil {
		return err
	}
	return os.RemoveAll(pathKey.FirstPathName())
}
func (s *Storage) read(key string) (io.Reader, error) {
	f, err := s.readStream(key)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, f)

	return buf, nil
}
func (s *Storage) readStream(key string) (io.ReadCloser, error) {
	pathKey := s.PathTransformFunc(key)

	return os.Open(pathKey.fullPath())

}
func (s *Storage) writeStream(key string, r io.Reader) error {
	pathKey := s.PathTransformFunc(key)
	if err := os.MkdirAll(pathKey.PathName, os.ModePerm); err != nil {
		return err
	}

	fullPath := pathKey.fullPath()

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	n, err := io.Copy(f, r)
	if err != nil {
		return err
	}

	log.Printf("written %d bytes to filename=%s\n", n, fullPath)
	return nil
}
func (p PathKey) fullPath() string {
	return fmt.Sprintf("%s/%s", p.PathName, p.FileName)
}
