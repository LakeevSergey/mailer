package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

type gzipReaderWriter struct {
	http.ResponseWriter
	Writer io.Writer
	Reader io.Reader
}

func (w gzipReaderWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (w gzipReaderWriter) Read(b []byte) (int, error) {
	return w.Reader.Read(b)
}

func Compressor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gzipReaderWriter{ResponseWriter: w, Writer: gz}, r)
	})
}

func Decompressor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Accept-Encoding", "gzip")

		if !strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewReader(r.Body)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		r.Body = gz
		next.ServeHTTP(w, r)
	})
}
