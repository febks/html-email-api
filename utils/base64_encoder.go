package utils

import (
	"encoding/base64"
	"io"
)

type base64LineWriter struct {
	w      io.Writer
	line   int
	buffer []byte
}

func newBase64LineWriter(w io.Writer) *base64LineWriter {
	return &base64LineWriter{w: w, buffer: make([]byte, 0, 76)}
}

func (lw *base64LineWriter) Write(p []byte) (int, error) {
	total := 0
	for len(p) > 0 {
		remain := 76 - lw.line
		n := len(p)
		if n > remain {
			n = remain
		}
		lw.buffer = append(lw.buffer, p[:n]...)
		lw.line += n
		p = p[n:]
		total += n

		if lw.line == 76 {
			if _, err := lw.w.Write(lw.buffer); err != nil {
				return total, err
			}
			if _, err := lw.w.Write([]byte("\r\n")); err != nil {
				return total, err
			}
			lw.buffer = lw.buffer[:0]
			lw.line = 0
		}
	}
	return total, nil
}

func (lw *base64LineWriter) Close() error {
	if len(lw.buffer) > 0 {
		if _, err := lw.w.Write(lw.buffer); err != nil {
			return err
		}
		if _, err := lw.w.Write([]byte("\r\n")); err != nil {
			return err
		}
	}
	return nil
}

func NewBase64Encoder(w io.Writer) io.WriteCloser {
	return base64.NewEncoder(base64.StdEncoding, newBase64LineWriter(w))
}
