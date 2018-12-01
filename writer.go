package textbuffer

import (
	"io"
	"sync"
)

// Writer maintains a buffer which is rotated every number of calls to
// Write method.
//
// Provides synchronized access to the underline writer
type Writer struct {
	maxWritesCnt int
	cnt          int
	buf          []byte
	bytesWritten int
	w            io.Writer

	mx sync.Mutex
}

// NewWriter creates an instance with corresponding writer and number of
// calls before the rotation.
func NewWriter(w io.Writer, cnt int) *Writer {
	return &Writer{maxWritesCnt: cnt, w: w}
}

// Write provided payload.
func (w *Writer) Write(p []byte) (int, error) {
	w.mx.Lock()
	defer w.mx.Unlock()

	w.buf = append(w.buf, p...)
	w.cnt++
	n := len(p)

	if w.cnt >= w.maxWritesCnt {
		return w.write()
	}

	w.bytesWritten += n

	return n, nil
}

// WriteString provided payload as a string.
func (w *Writer) WriteString(p string) (int, error) {
	return w.Write([]byte(p))
}

func (w *Writer) write() (int, error) {
	n, err := w.w.Write(w.buf)

	// there are less bytes accounted than it was written
	// but it needs to report only the diff submit with Write
	d := n - w.bytesWritten

	w.cnt = 0
	w.buf = nil
	w.bytesWritten = 0

	return d, err
}

// Flush writes buffered data into underline writer.
func (w *Writer) Flush() error {
	w.mx.Lock()
	defer w.mx.Unlock()

	_, err := w.write()
	return err
}
