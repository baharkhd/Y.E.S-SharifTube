package content

import "io"

// interface for Object Storage Driver
type OSDriver interface {
	Exists(path string) bool
	Store(path string, file io.Reader) error
	Get(path string)string
}
