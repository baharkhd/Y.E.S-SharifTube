package content

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sort"
	"time"
	"yes-sharifTube/graph/model"
	modelUtil "yes-sharifTube/internal/model"
	"yes-sharifTube/internal/model/comment"
)

type Content struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title        string             `json:"title" bson:"title"`
	Description  string             `json:"description" bson:"description"`
	Timestamp    int64              `json:"timestamp" bson:"timestamp"`
	UploadedByID string             `json:"uploaded_by" bson:"uploaded_by"`
	ApprovedByID string             `json:"approved_by" bson:"approved_by"`
	Vurl         string             `json:"vurl" bson:"vurl"` //todo better implementation
	Tags         []string           `json:"tags" bson:"tags"`
	Comments     []*comment.Comment `json:"comments" bson:"comments"`
	CourseID     string             `json:"course" bson:"course"`
}

func New(ID primitive.ObjectID, title, uploadedBy, vurl, courseID string, description, approvedBy *string, tags []string) (*Content, error) {
	err := RegexValidate(&title, description, &uploadedBy, &vurl, &courseID, approvedBy, tags)
	if err != nil {
		return nil, err
	}
	return &Content{
		ID:           ID,
		Title:        title,
		Description:  modelUtil.PtrTOStr(description),
		Timestamp:    time.Now().Unix(),
		UploadedByID: uploadedBy,
		ApprovedByID: modelUtil.PtrTOStr(approvedBy),
		Vurl:         vurl,
		Tags:         tags,
		Comments:     []*comment.Comment{},
		CourseID:     courseID,
	}, nil
}

func RegexValidate(title, description, uploadedBy, vurl, courseID, approvedBy *string, tags []string) error {
	//todo validate fields of a Content
	return nil
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
		return nil, model.InternalServerException{Message: "error while reshape comments of content: /n" + err.Error()}
	}
	res.Comments = comments

	return res, nil
}

func ReshapeAll(courses []*Content) ([]*model.Content, error) {
	var cs []*model.Content
	for _, c := range courses {
		tmp, err := c.Reshape()
		if err != nil {
			return nil, model.InternalServerException{Message: "error while reshape content array: /n" + err.Error()}
		}
		cs = append(cs, tmp)
	}
	return cs, nil
}

func Sort(arr []*Content) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Timestamp >= arr[j].Timestamp
	})
}

func GetAll(arr []*Content, start, amount int) []*Content {
	Sort(arr)
	n := len(arr)
	if start >= n {
		return []*Content{}
	}
	end := start + amount
	if end >= n {
		return arr[start:n]
	} else {
		return arr[start:end]
	}
}

func (c *Content) Update(newTitle, newDescription *string, newTags []string) error {
	if newTitle == nil && newDescription == nil {
		return model.EmptyFieldsException{Message: model.EmptyKeyErrorMessage}
	}
	err := RegexValidate(newTitle, newDescription, nil, nil, nil, nil, newTags)
	if err != nil {
		return err
	}
	if newTitle != nil {
		c.Title = *newTitle
	}
	if newDescription != nil {
		c.Description = *newDescription
	}
	if newTags != nil {
		c.Tags = newTags
	}
	c.Timestamp = time.Now().Unix()
	return nil
}

func (c *Content) GetComment(commentID primitive.ObjectID) (*comment.Comment, *comment.Reply) {
	for i, com := range c.Comments {
		if com.ID == commentID {
			return c.Comments[i], nil
		} else {
			for j, rep := range c.Comments[i].Replies {
				if rep.ID == commentID {
					return c.Comments[i], c.Comments[i].Replies[j]
				}
			}
		}
	}
	return nil, nil
}
