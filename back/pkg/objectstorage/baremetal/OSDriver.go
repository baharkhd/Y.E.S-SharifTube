package baremetal

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"path"
	"strings"
	"sync"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/pkg/objectstorage"
)

func (b BaremetalOSD) Exists(bucket *objectstorage.Bucket, filename string) bool {
	if err := b.Run(fmt.Sprintf("test -e %s/%s", bucket.GetPath(), filename)); err != nil {
		return false
	}
	return true
}

func (b BaremetalOSD) Store(bucket *objectstorage.Bucket, filename string, file io.Reader, size int64) error {
	if b.Exists(bucket, filename) {
		return model.FileAlreadyExistsException{Message: "file with the same name already exist!"}
	}

	return b.store(fmt.Sprintf("%s/%s", bucket.GetPath(), filename),file,size)
}

func (b *BaremetalOSD) StoreMulti(bucket *objectstorage.Bucket, filename string, files []*graphql.Upload) error {
	if b.Exists(bucket, filename) {
		return model.FileAlreadyExistsException{Message: "file with the same name already exist!"}
	}
	return b.storeMulti(fmt.Sprintf("%s/%s", bucket.GetPath(), filename),files)
}

func (b BaremetalOSD) Update(bucket *objectstorage.Bucket, filename string, file io.Reader, size int64) error {
	if b.Exists(bucket, filename) {
		if err := b.Run(fmt.Sprintf("rm -rf %s/%s", bucket.GetPath(), filename)); err != nil {
			return model.InternalServerException{Message: "couldn't update the file"}
		}
	}
	return b.store(fmt.Sprintf("%s/%s", bucket.GetPath(), filename),file,size)
}

func (b BaremetalOSD) GetURL(bucket *objectstorage.Bucket, filename string) string {
	hostname := strings.Split(b.host, ":")[0]
	path:= "/yes"+bucket.GetPath()[len(b.root.GetPath()):]
	return "http://"+strings.Replace(fmt.Sprintf("%s/%s/%s", hostname, path, filename),"//","/",1)
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


func (b *BaremetalOSD) storeMulti(pathInStorage string, files []*graphql.Upload) error {

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
		for _, file := range files {
			fmt.Fprintf(hostIn, "C0664 %d %s\n", file.Size, path.Base(pathInStorage))
			io.Copy(hostIn, file.File)
		}
		fmt.Fprint(hostIn, "\x00")
		wg.Done()
	}()

	session.Run(fmt.Sprintf("/usr/bin/scp -t %s",path.Dir(pathInStorage)))
	wg.Wait()
	return nil
}
