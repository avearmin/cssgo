package gomour

import (
	"io"
)

type Node interface {
	Render(w io.Writer) error
}

type statefulWriter struct {
	w   io.Writer
	err error
}

func (w *statefulWriter) Write(p []byte) {
	if w.err != nil {
		return
	}
	_, w.err = w.w.Write(p)
}
