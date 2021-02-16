package course

import "yes-sharifTube/pkg/objectstorage"

func (c * Course)getCourseBucket() *objectstorage.Bucket {
	for _, sub := range OSD.GetRoot().Subs {
		if sub.Id == c.ID.Hex() {
			return sub
		}
	}
	return OSD.NewBucket(OSD.GetRoot(),c.ID.Hex())
}
