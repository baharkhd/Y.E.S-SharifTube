package course

import "yes-sharifTube/pkg/objectstorage"

const attachmentPathBaseName = "attachments"

func (c * Course)getCourseBucket() *objectstorage.Bucket {
	for _, sub := range OSD.GetRoot().Subs {
		if sub.Id == c.ID.Hex() {
			return sub
		}
	}
	return OSD.NewBucket(OSD.GetRoot(),c.ID.Hex())
}

func (c *Course) GetAttachmentBucket() *objectstorage.Bucket {
	for _, sub := range c.getCourseBucket().Subs {
		if sub.Id == attachmentPathBaseName {
			return sub
		}
	}
	return OSD.NewBucket(c.getCourseBucket(),attachmentPathBaseName)
}

func Stream(vurl string) string{
	return OSD.Stream(vurl)
}