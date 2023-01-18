package dto

import "io"

type Add struct {
	FileName string
	Mime     string
	Content  io.ReadCloser
}
