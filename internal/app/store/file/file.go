package file

import (
	"io"
	"os"

	"github.com/FindHotel/emspy/internal/app/store"
)

type FileStore struct {
	file io.WriteCloser
}

func New(file string) (store.Store, error) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileStore{file: f}, nil
}

func (s *FileStore) InsertWebhook(source string, record interface{}) error {
	input := record.([]byte)

	_, err := s.file.Write(input)
	if err != nil {
		return err
	}
	_, err = s.file.Write([]byte("\n"))
	return err
}
