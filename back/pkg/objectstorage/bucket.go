package objectstorage

type Bucket struct {
	Id     string
	Parent *Bucket
	Subs   []*Bucket
}

func NewBucket(id string, parent *Bucket) *Bucket {
	b := &Bucket{
		Id:     id,
		Parent: parent,
		Subs:   []*Bucket{},
	}
	if parent != nil {
		parent.Subs = append(parent.Subs, b)
	}
	return b
}
func (b *Bucket)GetPath() string{
	if b.Parent !=nil{
		return b.Parent.GetPath()+"/"+b.Id
	}
	return b.Id
}