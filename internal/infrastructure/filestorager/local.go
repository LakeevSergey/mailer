package filestorager

import (
	"io"
	"os"
)

type LocalFileStorager struct {
	basePath          string
	filenameGenerator FilenameGeneratoor
}

func NewLocalFileStorager(basePath string, filenameGenerator FilenameGeneratoor) *LocalFileStorager {
	return &LocalFileStorager{
		basePath:          basePath,
		filenameGenerator: filenameGenerator,
	}
}

func (s *LocalFileStorager) Save(data io.Reader) (string, int, error) {
	bytes, err := io.ReadAll(data)
	if err != nil {
		return "", 0, err
	}

	filename := s.filenameGenerator()

	file, err := os.OpenFile(s.basePath+filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0777)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()

	_, err = file.Write(bytes)
	if err != nil {
		return "", 0, err
	}

	file.Close()
	if err != nil {
		return "", 0, err
	}

	return s.basePath + filename, len(bytes), nil
}

func (s *LocalFileStorager) Get(filename string) (io.ReadCloser, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *LocalFileStorager) Delete(filename string) error {
	return os.Remove(filename)
}
