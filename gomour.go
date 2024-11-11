package gomour

import (
	"io"
)

type Node interface {
	Render(w io.Writer) error
}
