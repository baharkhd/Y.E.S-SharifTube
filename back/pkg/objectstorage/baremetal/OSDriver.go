package baremetal

import (
	"fmt"
	"io"
	"path"
	"strings"
	"sync"
	"yes-sharifTube/graph/model"
)

func (b BaremetalObjectStorageDriver) Exists(path string) bool {
	if err := b.Run(fmt.Sprintf("test -e %s/%s", b.baseSir, path)); err != nil {
		return false
	}
	return true
}

func (b BaremetalObjectStorageDriver) Store(path string, file io.Reader, size int64) error {
	if b.Exists(path) {
		return model.FileAlreadyExistsException{Message: "file with the same name already exist!"}
	}


	return b.store(fmt.Sprintf("%s/%s", b.baseSir, path),file,size)
}

func (b BaremetalObjectStorageDriver) Update(path string, file io.Reader, size int64) error {
	if b.Exists(path) {
		if err := b.Run(fmt.Sprintf("rm -rf %s/%s", b.baseSir, path)); err != nil {
			return model.InternalServerException{Message: "couldn't update the file"}
		}
	}
	return b.store(fmt.Sprintf("%s/%s", b.baseSir, path),file,size)
}

func (b BaremetalObjectStorageDriver) GetURL(path string) string {
	hostname := strings.Split(b.host, ":")[0]
	return "http://"+strings.Replace(fmt.Sprintf("%s/%s/%s", hostname, b.baseSir, path),"//","/",1)
}

func (b BaremetalObjectStorageDriver) store(pathInStorage string, file io.Reader, size int64) error {

	wg := sync.WaitGroup{}
	wg.Add(1)

	session, err2 := b.client.NewSession()
	if err2 != nil {
		return err2
	}
	defer session.Close()

	go func() {
		hostIn, _ :=session.StdinPipe()
		defer hostIn.Close()
		fmt.Fprintf(hostIn, "C0664 %d %s\n", size, path.Base(pathInStorage))
		io.Copy(hostIn, file)
		fmt.Fprint(hostIn, "\x00")
		wg.Done()
	}()

	session.Run(fmt.Sprintf("/usr/bin/scp -t %s",path.Dir(pathInStorage)))
	wg.Wait()
	return nil
}
