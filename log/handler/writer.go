package handler

import (
	"bufio"
	"io"
	"os"
)

type WriteCloser struct {
	w   io.WriteCloser
	buf *bufio.Writer
}

func NewWriteCloser(w io.WriteCloser) *WriteCloser {
	return &WriteCloser{
		w:   w,
		buf: bufio.NewWriter(w),
	}
}

// Closed returns true if having been closed, or false.
func (wc *WriteCloser) Closed() bool {
	return wc.w == nil
}

// Write implements the interface io.Writer.
func (wc *WriteCloser) Write(data []byte) (int, error) {
	return wc.buf.Write(data)
}

// Close implements the interface io.Closer.
func (wc *WriteCloser) Close() (err error) {
	err = wc.w.Close()
	wc.w = nil
	wc.buf.Reset(os.Stderr)
	return err
}

// NullWriter is a null writer, which implements the interface io.WriteCloser.
// When writing the data, it will discard the data and return.
type NullWriter struct{}

// NewNullWriter returns a new null writer.
func NewNullWriter() NullWriter {
	return NullWriter{}
}

// Write writes the data to the handler.
func (w NullWriter) Write(d []byte) (int, error) {
	return len(d), nil
}

// WriteString writes the data to the handler.
func (w NullWriter) WriteString(d string) (n int, err error) {
	return len(d), nil
}

// Close close the handler.
func (w NullWriter) Close() (err error) {
	return nil
}
