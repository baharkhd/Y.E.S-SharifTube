package objectstorage

import (
	"github.com/99designs/gqlgen/graphql"
	"io"
)

// interface for Object Storage Driver
type OSDriver interface {
	Exists(bucket *Bucket, filename string) bool
	GetURL(bucket *Bucket, filename string) string
	Update(bucket *Bucket, filename string, file io.Reader, size int64) error
	Store(bucket *Bucket, filename string, file io.Reader, size int64) error
	StoreMulti(bucket *Bucket, filename string, files []*graphql.Upload) error
	GetRoot() *Bucket
	NewBucket(parent *Bucket, id string) *Bucket
	Stream(vurl string)
}
