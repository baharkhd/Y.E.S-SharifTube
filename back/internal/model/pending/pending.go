package pending

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sort"
	"time"
	"yes-sharifTube/graph/model"
	modelUtil "yes-sharifTube/internal/model"
)

const TitleWordSize = 30
const TitleCharSize = 150
const DescriptionWordSize = 200
const DescriptionCharSize = 600
const MessageWordSize = 200
const MessageCharSize = 600

type Pending struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title        string             `json:"title" bson:"title"`
	Description  string             `json:"description" bson:"description"`
	Status       Status             `json:"status" bson:"status"`
	Timestamp    int64              `json:"timestamp" bson:"timestamp"`
	UploadedByUn string             `json:"uploaded_by" bson:"uploaded_by"`
	Furl         string             `json:"furl" bson:"furl"` //todo better implementation
	CourseID     string             `json:"course" bson:"course"`
	Message      string             `json:"message" bson:"message"`
}

var DBD DBDriver

func New(title, uploadedByID, furl, courseID string, description *string) (*Pending, error) {
	err := RegexValidate(&title, description, &uploadedByID, &furl, &courseID, nil)
	if err != nil {
		return nil, err
	}
	p:= &Pending{
		Title:        title,
		Description:  modelUtil.PtrTOStr(description),
		Status:       PENDING,
		Timestamp:    time.Now().Unix(),
		UploadedByUn: uploadedByID,
		Furl:         furl,
		Message:      "",
		CourseID:     courseID,
	}
	// insert the pending into database
	pn, err := Insert(courseID, p)
	if err != nil {
		return nil, err
	}
	return pn,err

}

func RegexValidate(title, description, uploadedByID, furl, courseID, message *string) error {
	if title != nil && modelUtil.IsSTREmpty(*title) {
		return model.RegexMismatchException{Message: "title field is empty"}
	}
	if title != nil && (modelUtil.WordCount(*title) > TitleWordSize || len(*title) > TitleCharSize) {
		return model.RegexMismatchException{Message: "title field exceeds limit size"}
	}
	if description != nil && modelUtil.IsSTREmpty(*description) {
		return model.RegexMismatchException{Message: "description field is empty"}
	}
	if description != nil && (modelUtil.WordCount(*description) > DescriptionWordSize || len(*description) > DescriptionCharSize) {
		return model.RegexMismatchException{Message: "description field exceeds limit size"}
	}
	if uploadedByID != nil && modelUtil.IsSTREmpty(*uploadedByID) {
		return model.RegexMismatchException{Message: "uploader username field is empty"}
	}
	if message != nil && modelUtil.IsSTREmpty(*message) {
		return model.RegexMismatchException{Message: "message field is empty"}
	}
	if message != nil && (modelUtil.WordCount(*message) > MessageWordSize || len(*message) > MessageCharSize) {
		return model.RegexMismatchException{Message: "message field exceeds limit size"}
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

func (p *Pending) Update(newTitle, newDescription, message *string) error {
	if newTitle == nil && newDescription == nil {
		return model.EmptyFieldsException{Message: model.EmptyKeyErrorMessage}
	}
	err := RegexValidate(newTitle, newDescription, nil, nil, nil, message)
	if err != nil {
		return err
	}
	if newTitle != nil {
		p.Title = *newTitle
	}
	if newDescription != nil {
		p.Description = *newDescription
	}
	if message != nil {
		p.Message = *message
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
