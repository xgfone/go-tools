package handler

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
