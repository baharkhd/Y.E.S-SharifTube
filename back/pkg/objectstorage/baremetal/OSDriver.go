package baremetal

import (
	"fmt"
	"io"
	"path"
	"strings"
	"sync"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/pkg/objectstorage"
)

func (b BaremetalOSD) Exists(bucket *objectstorage.Bucket,path string) bool {
	if err := b.Run(fmt.Sprintf("test -e %s/%s", b.root.GetPath(), path)); err != nil {
		return false
	}
	return true
}

func (b BaremetalOSD) Store(bucket *objectstorage.Bucket,path string, file io.Reader, size int64) error {
	if b.Exists(bucket,path) {
		return model.FileAlreadyExistsException{Message: "file with the same name already exist!"}
	}


	return b.store(fmt.Sprintf("%s/%s", b.root.GetPath(), path),file,size)
}

func (b BaremetalOSD) Update(bucket *objectstorage.Bucket,path string, file io.Reader, size int64) error {
	if b.Exists(bucket,path) {
		if err := b.Run(fmt.Sprintf("rm -rf %s/%s", b.root.GetPath(), path)); err != nil {
			return model.InternalServerException{Message: "couldn't update the file"}
		}
	}
	return b.store(fmt.Sprintf("%s/%s", b.root.GetPath(), path),file,size)
}

func (b BaremetalOSD) GetURL(bucket *objectstorage.Bucket,path string) string {
	hostname := strings.Split(b.host, ":")[0]
	return "http://"+strings.Replace(fmt.Sprintf("%s/%s/%s", hostname, b.root.GetPath(), path),"//","/",1)
}

func (b BaremetalOSD) store(pathInStorage string, file io.Reader, size int64) error {

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
