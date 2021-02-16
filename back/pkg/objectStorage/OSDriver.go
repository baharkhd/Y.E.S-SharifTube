package objectStorage

import "io"

// interface for Object Storage Driver
type OSDriver interface {
	Exists(path string) bool
	Store(path string, file io.Reader) error
	Update(path string, file io.Reader) error
	GetURL(path string) string
}
