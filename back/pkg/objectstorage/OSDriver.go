package objectstorage

import "io"

// interface for Object Storage Driver
type OSDriver interface {
	Exists(bucket *Bucket, path string) bool
	GetURL(bucket *Bucket, path string) string
	Update(bucket *Bucket, path string, file io.Reader, size int64) error
	Store(bucket *Bucket, path string, file io.Reader, size int64) error
	GetRoot() *Bucket
	NewBucket(parent *Bucket, id string) *Bucket
}
