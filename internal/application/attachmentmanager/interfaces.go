package attachmentmanager

import "io"

type FileStorager interface {
	Save(data io.Reader) (string, int, error)
	Get(name string) (io.ReadCloser, error)
	Delete(filename string) error
}
