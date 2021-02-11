package attachment

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"yes-sharifTube/graph/model"
	modelUtil "yes-sharifTube/internal/model"
)

type Attachment struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Timestamp   int64              `json:"timestamp" bson:"timestamp"`
	Aurl        string             `json:"aurl" bson:"aurl"` //todo better implementation
	CourseID    string             `json:"course" bson:"course"`
}

func New(ID primitive.ObjectID, name, aurl, courseID string, description *string) (*Attachment, error) {
	err := RegexValidate(&name, description, &aurl, &courseID)
	if err != nil {
		return nil, err
	}
	return &Attachment{
		ID:          ID,
		Name:        name,
		Description: modelUtil.PtrTOStr(description),
		Timestamp:   time.Now().Unix(),
		Aurl:        aurl,
		CourseID:    courseID,
	}, nil
}

func RegexValidate(name, description, aurl, courseID *string) error {
	if name != nil && modelUtil.IsSTREmpty(*name) {
		return model.RegexMismatchException{Message: "name field is empty"}
	}
	if description != nil && modelUtil.IsSTREmpty(*description) {
		return model.RegexMismatchException{Message: "description field is empty"}
	}
	//todo regex definition for Aurl field
	if aurl != nil && modelUtil.IsSTREmpty(*aurl) {
		return model.RegexMismatchException{Message: "file URL is empty"}
	}
	if courseID != nil {
		_, err := primitive.ObjectIDFromHex(*courseID)
		if err != nil {
			return model.RegexMismatchException{Message: "courseID field is invalid"}
		}
	}
	return nil
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

func (a *Attachment) Update(newName, newDescription *string) error {
	if newName == nil && newDescription == nil {
		return model.EmptyFieldsException{Message: model.EmptyKeyErrorMessage}
	}
	err := RegexValidate(newName, newDescription, nil, nil)
	if err != nil {
		return err
	}
	if newName != nil {
		a.Name = *newName
	}
	if newDescription != nil {
		a.Description = *newDescription
	}
	a.Timestamp = time.Now().Unix()
	return nil
}
