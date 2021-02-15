package content

import (
	"encoding/json"
	"github.com/coocood/freecache"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sort"
	"time"
	"yes-sharifTube/graph/model"
	modelUtil "yes-sharifTube/internal/model"
	"yes-sharifTube/internal/model/comment"
)

const CacheExpire = 10 * 60
const TitleWordSize = 30
const TitleCharSize = 200
const DescriptionWordSize = 300
const DescriptionCharSize = 1200

type Content struct {
	ID           primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title        string             `json:"title" bson:"title"`
	Description  string             `json:"description" bson:"description"`
	Timestamp    int64              `json:"timestamp" bson:"timestamp"`
	UploadedByUn string             `json:"uploaded_by" bson:"uploaded_by"`
	ApprovedByUn string             `json:"approved_by" bson:"approved_by"`
	Vurl         string             `json:"vurl" bson:"vurl"` //todo better implementation
	Tags         []string           `json:"tags" bson:"tags"`
	Comments     []*comment.Comment `json:"comments" bson:"comments"`
	CourseID     string             `json:"course" bson:"course"`
}

var DBD DBDriver
var Cache *freecache.Cache

func New(title, uploadedBy, vurl, courseID string, description, approvedBy *string, tags []string) (*Content, error) {
	err := RegexValidate(&title, description, &uploadedBy, &vurl, &courseID, approvedBy, tags)
	if err != nil {
		return nil, err
	}
	return &Content{
		Title:        title,
		Description:  modelUtil.PtrTOStr(description),
		Timestamp:    time.Now().Unix(),
		UploadedByUn: uploadedBy,
		ApprovedByUn: modelUtil.PtrTOStr(approvedBy),
		Vurl:         vurl,
		Tags:         tags,
		Comments:     []*comment.Comment{},
		CourseID:     courseID,
	}, nil
}

func RegexValidate(title, description, uploadedBy, vurl, courseID, approvedBy *string, tags []string) error {
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
	if uploadedBy != nil && modelUtil.IsSTREmpty(*uploadedBy) {
		return model.RegexMismatchException{Message: "uploader username field is empty"}
	}
	if approvedBy != nil && modelUtil.IsSTREmpty(*approvedBy) {
		return model.RegexMismatchException{Message: "approves  username field is empty"}
	}
	//todo regex definition for Vurl field
	if vurl != nil && modelUtil.IsSTREmpty(*vurl) {
		return model.RegexMismatchException{Message: "file URL is empty"}
	}
	if courseID != nil {
		_, err := primitive.ObjectIDFromHex(*courseID)
		if err != nil {
			return model.RegexMismatchException{Message: "courseID field is invalid"}
		}
	}
	//todo regex definition for Tag
	if tags != nil && len(tags) == 0{
		return model.RegexMismatchException{Message: "tags field is empty"}
	}
	return nil
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

func GetFromCache(contentID string) (*Content, error){
	c, err := Cache.Get([]byte(contentID))
	if err == nil {
		var cr *Content
		err = json.Unmarshal(c, &cr)
		if err != nil {
			return nil, model.InternalServerException{Message: err.Error()}
		}
		return cr, err
	}
	return nil,  model.ContentNotFoundException{Message: "content not found in cache"}
}

func (c *Content) Cache() error {
	content, err := json.Marshal(c)
	if err != nil {
		return model.InternalServerException{Message: err.Error()}
	}
	err = Cache.Set([]byte(c.ID.Hex()), content, CacheExpire)
	if err != nil {
		return model.InternalServerException{Message: "content couldn't cache"}
	}
	return nil
}