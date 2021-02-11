package course

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"yes-sharifTube/graph/model"
	modelUtils "yes-sharifTube/internal/model"
	"yes-sharifTube/internal/model/attachment"
	"yes-sharifTube/internal/model/comment"
	"yes-sharifTube/internal/model/content"
	"yes-sharifTube/internal/model/pending"
)

type Course struct {
	ID        primitive.ObjectID       `bson:"_id" json:"id,omitempty"`
	Title     string                   `json:"title" bson:"title"`
	Summery   string                   `json:"summery" bson:"summery"`
	CreatedAt int64                    `json:"created_at" bson:"created_at"`
	Token     string                   `json:"token" bson:"token"`
	ProfUn    string                   `json:"prof" bson:"prof"`
	TaUns     []string                 `json:"tas" bson:"tas"`
	StdUns    []string                 `json:"students" bson:"students"`
	Contents  []*content.Content       `json:"contents" bson:"contents"`
	Pends     []*pending.Pending       `json:"pends" bson:"pends"`
	Inventory []*attachment.Attachment `json:"inventory" bson:"inventory"`
}

func New(title, summery, profUsername, token string) (*Course, error) {
	hashedToken, err := modelUtils.HashToken([]byte(token))
	if err != nil {
		return nil, model.InternalServerException{Message: "internal server error: couldn't hash token"}
	}

	return &Course{
		Title:     title,
		Summery:   summery,
		CreatedAt: time.Now().Unix(),
		ProfUn:    profUsername,
		Token:     hashedToken,
		TaUns:     []string{},
		StdUns:    []string{},
		Contents:  []*content.Content{},
		Pends:     []*pending.Pending{},
		Inventory: []*attachment.Attachment{},
	}, nil
}

func (c Course) Reshape() (*model.Course, error) {
	//todo get Users from database by usernames
	var prof *model.User
	var tas []*model.User
	var students []*model.User

	res := &model.Course{
		ID:        c.ID.Hex(),
		Title:     c.Title,
		Summary:   &c.Summery,
		CreatedAt: int(c.CreatedAt),
		Prof:      prof,
		Tas:       tas,
		Pends:     nil,
		Students:  students,
		Contents:  nil,
		Inventory: nil,
	}

	//reshape pendings
	pends, err := pending.ReshapeAll(c.Pends)
	if err != nil {
		return nil, &model.InternalServerException{Message: "error while reshape pending array of course: /n" + err.Error()}
	}
	res.Pends = pends

	//reshape contents
	contents, err := content.ReshapeAll(c.Contents)
	if err != nil {
		return nil, &model.InternalServerException{Message: "error while reshape contents of course: /n" + err.Error()}
	}
	res.Contents = contents

	//reshape inventory
	res.Inventory = attachment.ReshapeAll(c.Inventory)

	return res, nil
}

func ReshapeAll(courses []*Course) ([]*model.Course, error) {
	var cs []*model.Course
	for _, c := range courses {
		tmp, err := c.Reshape()
		if err != nil {
			return nil, &model.InternalServerException{Message: "error while reshape course array: " + err.Error()}
		}
		cs = append(cs, tmp)
	}
	return cs, nil
}

func (c *Course) UpdateInfo(newTitle, newSummery, newToken string) {
	c.Title = newTitle
	c.Summery = newSummery
	c.Token = newToken
}

func (c Course) IsUserNotStudent(username string) bool {
	if c.ProfUn == username || modelUtils.ContainsInStringArray(c.TaUns, username) {
		return true
	}
	return false
}

func (c Course) IsUserProfessor(username string) bool {
	return c.ProfUn == username
}

func (c Course) IsUserStudent(username string) bool {
	return modelUtils.ContainsInStringArray(c.StdUns, username)
}

func (c Course) IsUserTA(username string) bool {
	return modelUtils.ContainsInStringArray(c.TaUns, username)
}

func (c Course) IsUserAllowedToDeleteUser(username, target string) bool {
	// professor can remove every one except his self
	if c.ProfUn == username && username != target {
		return true
	}
	// ta can remove every one except professor
	if modelUtils.ContainsInStringArray(c.TaUns, username) && c.ProfUn != target {
		return true
	}
	// every body can remove them selves except professor
	if modelUtils.ContainsInStringArray(c.StdUns, username) && username == target && c.ProfUn != username {
		return true
	}
	return false
}

func (c Course) IsUserAllowedToDeleteUserComment(username, target string) bool {
	// professor can remove every one except his self
	if c.ProfUn == username {
		return true
	}
	// ta can remove every one except professor
	if modelUtils.ContainsInStringArray(c.TaUns, username) && c.ProfUn != target {
		return true
	}
	// every body can remove them selves except professor
	if modelUtils.ContainsInStringArray(c.StdUns, username) && username == target {
		return true
	}
	return false
}

func (c Course) IsUserParticipateInCourse(username string) bool {
	if c.ProfUn == username || modelUtils.ContainsInStringArray(c.TaUns, username) || modelUtils.ContainsInStringArray(c.StdUns, username) {
		return true
	}
	return false
}

func (c Course) CheckCourseToken(token string) bool {
	return modelUtils.CheckTokenHash(token, c.Token)
}

func (c Course) GetContent(contentID primitive.ObjectID) *content.Content {
	for _, cnt := range c.Contents {
		if cnt.ID == contentID {
			return cnt
		}
	}
	return nil
}

func (c Course) GetPending(contentID primitive.ObjectID) *pending.Pending {
	for _, pnt := range c.Pends {
		if pnt.ID == contentID {
			return pnt
		}
	}
	return nil
}

func (c Course) GetAttachment(attachmentID primitive.ObjectID) *attachment.Attachment {
	for _, pnt := range c.Inventory {
		if pnt.ID == attachmentID {
			return pnt
		}
	}
	return nil
}

func (c *Course) RemoveComment(username string, commentID primitive.ObjectID, cnt *content.Content) (*comment.Comment, *comment.Reply, error) {
	for i, com := range cnt.Comments {
		if com.ID == commentID {
			if !c.IsUserAllowedToDeleteUserComment(username, com.AuthorUn) {
				return nil, nil, &model.UserNotAllowedException{Message: "NotAllowed"}
			}
			rc := cnt.Comments[i]
			cnt.Comments = append(cnt.Comments[:i], cnt.Comments[i+1:]...)
			return rc, nil, nil
		} else {
			for j, rep := range cnt.Comments[i].Replies {
				if rep.ID == commentID {
					if !c.IsUserAllowedToDeleteUserComment(username, rep.AuthorUn) {
						return nil, nil, &model.UserNotAllowedException{Message: "NotAllowed"}
					}
					rc := cnt.Comments[i].Replies[j]
					cnt.Comments[i].Replies = append(cnt.Comments[i].Replies[:j], cnt.Comments[i].Replies[j+1:]...)
					return nil, rc, nil
				}
			}
		}
	}
	return nil, nil, &model.CommentNotFoundException{Message: "NotFound"}
}
