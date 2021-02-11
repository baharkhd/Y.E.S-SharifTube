package pending

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sort"
	"time"
	"yes-sharifTube/graph/model"
	modelUtil "yes-sharifTube/internal/model"
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

func New(ID primitive.ObjectID, title, uploadedByID, furl, courseID string, description *string) (*Pending, error) {
	err := RegexValidate(&title, description, &uploadedByID, &furl, &courseID)
	if err != nil {
		return nil, err
	}
	return &Pending{
		ID:           ID,
		Title:        title,
		Description:  modelUtil.PtrTOStr(description),
		Status:       PENDING,
		Timestamp:    time.Now().Unix(),
		UploadedByUn: uploadedByID,
		Furl:         furl,
		CourseID:     courseID,
	}, nil
}

func RegexValidate(title, description, uploadedByID, furl, courseID *string) error {
	if title != nil && modelUtil.IsSTREmpty(*title) {
		return model.RegexMismatchException{Message: "title field is empty"}
	}
	if description != nil && modelUtil.IsSTREmpty(*description) {
		return model.RegexMismatchException{Message: "description field is empty"}
	}
	if uploadedByID != nil && modelUtil.IsSTREmpty(*uploadedByID) {
		return model.RegexMismatchException{Message: "uploader username field is empty"}
	}
	//todo regex definition for Furl field
	if furl != nil && modelUtil.IsSTREmpty(*furl) {
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
			return nil, model.InternalServerException{Message: "error while reshape pending array: /n" + err.Error()}
		}
		ps = append(ps, tmp)
	}
	return ps, nil
}

func (p *Pending) Update(newTitle, newDescription *string) error {
	if newTitle == nil && newDescription == nil {
		return model.EmptyFieldsException{Message: model.EmptyKeyErrorMessage}
	}
	err := RegexValidate(newTitle, newDescription, nil, nil, nil)
	if err != nil {
		return err
	}
	if newTitle != nil {
		p.Title = *newTitle
	}
	if newDescription != nil {
		p.Description = *newDescription
	}
	p.Timestamp = time.Now().Unix()
	return nil
}

func (p *Pending) Accept() {
	p.Status = ACCEPTED
}

func (p *Pending) Reject() {
	p.Status = REJECTED
}

func Sort(arr []*Pending) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Timestamp >= arr[j].Timestamp
	})
}

func GetAll(arr []*Pending, start, amount int) []*Pending {
	Sort(arr)
	n := len(arr)
	if start >= n {
		return []*Pending{}
	}
	end := start + amount
	if end >= n {
		return arr[start:n]
	} else {
		return arr[start:end]
	}
}
