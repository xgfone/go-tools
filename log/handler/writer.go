package handler

// A null writer, which implements the interface io.WriteCloser. When writing
// the data, it will discard the data and return.
type NullWriter struct{}

// Get a new null writer.
func NewNullWriter() NullWriter {
	return NullWriter{}
}

func (w NullWriter) Write(d []byte) (int, error) {
	return len(d), nil
}

func (w NullWriter) WriteString(d string) (n int, err error) {
	return len(d), nil
}

func (w NullWriter) Close() (err error) {
	return nil
}
