package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	return s

}

func teardown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}
func TestPathTransformFunc(t *testing.T) {
	key := "vacationpictures"
	pathKey := CASPathTransformFunc(key)
	expectedFilename := "cd55ed3cc7b673af81e275b1fb8caa738b0b680e"
	expectedPathName := "cd55e/d3cc7/b673a/f81e2/75b1f/b8caa/738b0/b680e"
	if pathKey.PathName != expectedPathName {
		t.Errorf("have %s want %s", pathKey.PathName, expectedPathName)
	}
	if pathKey.Filename != expectedFilename {
		t.Errorf("have %s want %s", pathKey.Filename, expectedFilename)
	}
}

func TestDeleteStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "myspecialpicture"
	data := []byte("some jpeg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	if err := s.Delete(key); err != nil {
		t.Error(err)
	}
	if s.Has(key) {
		t.Error("expected false got true")
	}
}
func TestStore(t *testing.T) {
	s := newStore()

	defer teardown(t, s)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("foo_%d", i)
		data := []byte("some jpeg bytes")
		if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}
		if ok := s.Has(key); !ok {
			t.Errorf("expected to have key %s", key)
		}
		r, err := s.Read(key)
		if err != nil {
			t.Error(err)
		}
		b, _ := io.ReadAll(r)
		if string(b) != string(data) {
			t.Errorf("want %s got %s", data, b)
		}
		fmt.Println(string(b))
		s.Delete(key)
		if ok := s.Has(key); ok {
			t.Errorf("did not expect to have key %s", key)
		}
	}

}
