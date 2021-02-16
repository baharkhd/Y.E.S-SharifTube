package baremetal

import "io"

func (b BaremetalObjectStorageDriver) Exists(path string) bool {
	panic("implement me")
}

func (b BaremetalObjectStorageDriver) Store(path string, file io.Reader) error {
	panic("implement me")
}

func (b BaremetalObjectStorageDriver) Update(path string, file io.Reader) error {
	panic("implement me")
}

func (b BaremetalObjectStorageDriver) GetURL(path string) string {
	panic("implement me")
}

