package baremetal

import (
	"golang.org/x/crypto/ssh"
	"yes-sharifTube/pkg/objectstorage"
)

var publicKeyPath = "/tmp/id_rsa"

type BaremetalOSD struct {
	root   *objectstorage.Bucket
	host   string
	client *ssh.Client
}


func (b *BaremetalOSD) GetRoot() *objectstorage.Bucket {
	return b.root
}

func (b *BaremetalOSD) NewBucket(parent *objectstorage.Bucket, id string) *objectstorage.Bucket {
	bucket := objectstorage.NewBucket(id, parent)
	b.Run("/usr/bin/mkdir -p " + bucket.GetPath())
	return bucket
}

func New(host, username, baseDir string) (*BaremetalOSD, error) {
	client, err := newSSHClient(host, username, publicKeyPath)
	if err != nil {
		return nil, err
	}

	return &BaremetalOSD{
		root:   objectstorage.NewBucket(baseDir, nil),
		host:   host,
		client: client,
	}, nil
}
