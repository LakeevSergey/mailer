package entity

import "io"

type File struct {
	Info FileInfo
	Data io.ReadCloser
}
