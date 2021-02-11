package attachment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"yes-sharifTube/graph/model"
)

type Attachment struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Timestamp   int64              `json:"timestamp" bson:"timestamp"`
	Aurl        string             `json:"aurl" bson:"aurl"` //todo better implementation
	CourseID    string             `json:"course" bson:"course"`
}

func New(name, description, aurl, courseID string) (*Attachment, error) {
	//todo regex checking for url field

	return &Attachment{
		Name:        name,
		Description: description,
		Timestamp:   time.Now().Unix(),
		Aurl:        aurl,
		CourseID:    courseID,
	}, nil
}

func (a Attachment) Reshape() *model.Attachment {
	return &model.Attachment{
		ID:          a.ID.Hex(),
		Name:        a.Name,
		Aurl:        a.Aurl,
		Description: &a.Description,
		Timestamp:   int(a.Timestamp),
		CourseID:    a.CourseID,
	}
}

func ReshapeAll(courses []*Attachment) []*model.Attachment {
	var cs []*model.Attachment
	for _, c := range courses {
		cs = append(cs, c.Reshape())
	}
	return cs
}

func (a *Attachment) Update(newName, newDescription string) {
	a.Name = newName
	a.Description = newDescription
	a.Timestamp = time.Now().Unix()
}
