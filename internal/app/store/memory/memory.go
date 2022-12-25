package memory

import (
	"bufio"
	"bytes"
	"io"

	"github.com/FindHotel/emspy/internal/app/store"
)

type MemoryStore struct {
	memory io.ReadWriter
}

func New() (store.Store, error) {
	b := bytes.Buffer{}
	w := bufio.NewReadWriter(bufio.NewReader(&b), bufio.NewWriter(&b))
	return &MemoryStore{memory: w}, nil
}

func (s *MemoryStore) InsertWebhook(source string, record interface{}) error {
	input := record.([]byte)

	_, err := s.memory.Write(input)
	if err != nil {
		return err
	}
	_, err = s.memory.Write([]byte("\n"))
	return err
}

func (s *MemoryStore) Dump() ([]byte, error) {
	return io.ReadAll(s.memory)
}
