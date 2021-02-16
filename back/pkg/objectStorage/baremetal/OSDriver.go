package baremetal

import (
	"fmt"
	"io"
	"yes-sharifTube/graph/model"
)

func (b BaremetalObjectStorageDriver) Exists(path string) bool {
	if err := b.Run(fmt.Sprintf("test -e %s/%s", b.baseSir, path)); err != nil {
		return false
	}
	return true
}

func (b BaremetalObjectStorageDriver) Store(path string, file io.Reader) error {
	if b.Exists(path) {
		return model.FileAlreadyExistsException{Message: "file with the same name already exist!"}
	}
	return b.scpSession.CopyFile(file, fmt.Sprintf("%s/%s", b.baseSir, path), "777")
}

func (b BaremetalObjectStorageDriver) Update(path string, file io.Reader) error {
	if b.Exists(path) {
		if err := b.Run(fmt.Sprintf("rm -rf %s/%s", b.baseSir, path)); err != nil {
			return model.InternalServerException{Message: "couldn't update the file"}
		}
	}
	return b.scpSession.CopyFile(file, fmt.Sprintf("%s/%s", b.baseSir, path), "777")
}

func (b BaremetalObjectStorageDriver) GetURL(path string) string {
	return fmt.Sprintf("http://%s/%s/%s", b.host, b.baseSir, path)
}
