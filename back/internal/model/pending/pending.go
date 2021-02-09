package pending

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"yes-sharifTube/graph/model"
)

type Pending struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title        string             `json:"title" bson:"title"`
	Description  string             `json:"description" bson:"description"`
	Status       Status             `json:"status" bson:"status"`
	Timestamp    int64              `json:"timestamp" bson:"timestamp"`
	UploadedByUn string             `json:"uploaded_by" bson:"uploaded_by"`
	Furl         string             `json:"furl" bson:"furl"` //todo better implementation
	CourseID     string             `json:"course" bson:"course"`
}

func New(title, description, uploadedByID, furl, courseID string) (*Pending, error) {
	//todo regex checking for url field

	return &Pending{
		Title:        title,
		Description:  description,
		Status:       PENDING,
		Timestamp:    time.Now().Unix(),
		UploadedByUn: uploadedByID,
		Furl:         furl,
		CourseID:     courseID,
	}, nil
}

func (p Pending) Reshape() (*model.Pending, error) {
	//todo get author user by its username from database
	var uploader *model.User

	return &model.Pending{
		ID:          p.ID.Hex(),
		Title:       p.Title,
		Description: &p.Description,
		Status:      p.Status.Reshape(),
		Timestamp:   int(p.Timestamp),
		UploadedBy:  uploader,
		Furl:        p.Furl,
		CourseID:    p.CourseID,
	}, nil
}

func ReshapeAll(pendings []*Pending) ([]*model.Pending, error) {
	var ps []*model.Pending
	for _, p := range pendings {
		tmp, err := p.Reshape()
		if err != nil {
			return nil, &model.InternalServerException{Message: "error while reshape pending array: /n" + err.Error()}
		}
		ps = append(ps, tmp)
	}
	return ps, nil
}
