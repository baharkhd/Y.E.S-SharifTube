package content

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/model/comment"
)

type Content struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title        string             `json:"title" bson:"title"`
	Description  string             `json:"description" bson:"description"`
	Timestamp    int64              `json:"timestamp" bson:"timestamp"`
	UploadedByID string             `json:"uploaded_by" bson:"uploaded_by"`
	ApprovedByID *string            `json:"approved_by" bson:"approved_by"`
	Vurl         string             `json:"vurl" bson:"vurl"` //todo better implementation
	Tags         []string           `json:"tags" bson:"tags"`
	Comments     []*comment.Comment `json:"comments" bson:"comments"`
	CourseID     string             `json:"course" bson:"course"`
}

func New(title, description, uploadedBy, vurl, courseID string, approvedBy *string, tags []string) (*Content, error) {
	//todo regex checking for url field

	return &Content{
		Title:        title,
		Description:  description,
		Timestamp:    time.Now().Unix(),
		UploadedByID: uploadedBy,
		ApprovedByID: approvedBy,
		Vurl:         vurl,
		Tags:         tags,
		Comments:     []*comment.Comment{},
		CourseID:     courseID,
	}, nil
}

func (c Content) Reshape() (*model.Content, error) {
	//todo get uploadedBy & approvedBy users from database
	var uploader *model.User
	var approver *model.User

	res := &model.Content{
		ID:          c.ID.Hex(),
		Title:       c.Title,
		Description: &c.Description,
		Timestamp:   int(c.Timestamp),
		UploadedBy:  uploader,
		ApprovedBy:  approver,
		Vurl:        c.Vurl,
		Tags:        c.Tags,
		Comments:    nil,
		CourseID:    c.CourseID,
	}

	//reshape comments
	comments, err := comment.ReshapeAll(c.Comments)
	if err != nil {
		return nil, &model.InternalServerException{Message: "error while reshape comments of content: /n" + err.Error()}
	}
	res.Comments = comments

	return res, nil
}

func ReshapeAll(courses []*Content) ([]*model.Content, error) {
	var cs []*model.Content
	for _, c := range courses {
		tmp, err := c.Reshape()
		if err != nil {
			return nil, &model.InternalServerException{Message: "error while reshape content array: /n" + err.Error()}
		}
		cs = append(cs, tmp)
	}
	return cs, nil
}
