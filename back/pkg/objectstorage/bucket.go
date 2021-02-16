package objectstorage

type Bucket struct {
	id     string
	parent *Bucket
	subs   []*Bucket
}

func NewBucket(id string, parent *Bucket) *Bucket {
	b := &Bucket{
		id:     id,
		parent: parent,
		subs:   []*Bucket{},
	}
	if parent != nil {
		parent.subs = append(parent.subs, b)
	}
	return b
}
func (b *Bucket)GetPath() string{
	if b.parent!=nil{
		return b.parent.GetPath()+"/"+b.id
	}
	return b.id
}