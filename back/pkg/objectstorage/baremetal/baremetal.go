package baremetal

import (
	"golang.org/x/crypto/ssh"
	"io"
	"yes-sharifTube/pkg/objectstorage"
)

var publicKeyPath = "/home/kycilius/.ssh/id_rsa"

type BaremetalObjectStorageDriver struct {
	root   *objectstorage.Bucket
	host   string
	client *ssh.Client
}

func (b *BaremetalObjectStorageDriver) Exists(bucket *objectstorage.Bucket, path string) bool {
	panic("implement me")
}

func (b *BaremetalObjectStorageDriver) GetURL(bucket *objectstorage.Bucket, path string) string {
	panic("implement me")
}

func (b *BaremetalObjectStorageDriver) Update(bucket *objectstorage.Bucket, path string, file io.Reader, size int64) error {
	panic("implement me")
}

func (b *BaremetalObjectStorageDriver) Store(bucket *objectstorage.Bucket, path string, file io.Reader, size int64) error {
	panic("implement me")
}

func (b *BaremetalObjectStorageDriver) GetRoot() *objectstorage.Bucket {
	return b.root
}

func (b *BaremetalObjectStorageDriver) NewBucket(parent *objectstorage.Bucket, id string) *objectstorage.Bucket {
	bucket := objectstorage.NewBucket(id, parent)

}

func New(host, username, baseDir string) (*BaremetalObjectStorageDriver, error) {
	client, err := newSSHClient(host, username, publicKeyPath)
	if err != nil {
		return nil, err
	}

	return &BaremetalObjectStorageDriver{
		root:   objectstorage.NewBucket(baseDir, nil),
		host:   host,
		client: client,
	}, nil
}
