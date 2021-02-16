package course

import (
	"encoding/json"
	"github.com/99designs/gqlgen/graphql"
	"github.com/coocood/freecache"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
	"yes-sharifTube/graph/model"
	modelUtil "yes-sharifTube/internal/model"
	"yes-sharifTube/internal/model/attachment"
	"yes-sharifTube/internal/model/comment"
	"yes-sharifTube/internal/model/content"
	"yes-sharifTube/internal/model/pending"
	"yes-sharifTube/pkg/objectstorage"
)

const CacheExpire = 10 * 60
const TitleWordSize = 30
const TitleCharSize = 150
const SummeryWordSize = 200
const SummeryCharSize = 800

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

var (
	Cache *freecache.Cache
	OSD   objectstorage.OSDriver
	DBD   DBDriver
)

func New(title, profUsername, token string, summery *string) (*Course, error) {
	hashedToken, err := modelUtil.HashToken([]byte(token))
	if err != nil {
		return nil, model.InternalServerException{Message: "internal server error: couldn't hash token"}
	}
	err = RegexValidate(&title, summery, &profUsername, &token)
	if err != nil {
		return nil, err
	}
	return &Course{
		Title:     title,
		Summery:   modelUtil.PtrTOStr(summery),
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

func RegexValidate(title, summery, profUsername, token *string) error {
	if title != nil && modelUtil.IsSTREmpty(*title) {
		return model.RegexMismatchException{Message: "title field is empty"}
	}
	if title != nil && (modelUtil.WordCount(*title) > TitleWordSize || len(*title) > TitleCharSize) {
		return model.RegexMismatchException{Message: "title field exceeds limit size"}
	}
	if summery != nil && modelUtil.IsSTREmpty(*summery) {
		return model.RegexMismatchException{Message: "summery field is empty"}
	}
	if summery != nil && (modelUtil.WordCount(*summery) > SummeryWordSize || len(*summery) > SummeryCharSize) {
		return model.RegexMismatchException{Message: "summery field exceeds limit size"}
	}
	if profUsername != nil && modelUtil.IsSTREmpty(*profUsername) {
		return model.RegexMismatchException{Message: "professor username field is empty"}
	}
	//todo regex definition for token field
	if token != nil && modelUtil.IsSTREmpty(*token) {
		return model.RegexMismatchException{Message: "file URL is empty"}
	}
	return nil
}

func (c *Course) Update(newTitle, newSummery, newToken *string) error {
	if newTitle == nil && newSummery == nil && newToken == nil {
		return model.EmptyFieldsException{Message: model.EmptyKeyErrorMessage}
	}
	err := RegexValidate(newTitle, newSummery, nil, nil)
	if err != nil {
		return err
	}
	if newTitle != nil {
		c.Title = *newTitle
	}
	if newSummery != nil {
		c.Summery = *newSummery
	}
	if newToken != nil {
		hashedToken, err := modelUtil.HashToken([]byte(*newToken))
		if err != nil {
			return model.InternalServerException{Message: "internal server error: couldn't hash token"}
		}
		c.Token = hashedToken
	}
	return nil
}

func GetFromCache(courseID string) (*Course, error) {
	c, err := Cache.Get([]byte(courseID))
	if err == nil {
		var cr *Course
		err = json.Unmarshal(c, &cr)
		if err != nil {
			return nil, model.InternalServerException{Message: err.Error()}
		}
		return cr, err
	}
	return nil, model.CourseNotFoundException{Message: "course not found in cache"}
}

func (c *Course) Cache() error {
	course, err := json.Marshal(c)
	if err != nil {
		return model.InternalServerException{Message: err.Error()}
	}
	err = Cache.Set([]byte(c.ID.Hex()), course, CacheExpire)
	if err != nil {
		return model.InternalServerException{Message: "course couldn't in cache"}
	}
	return nil
}

func (c *Course) UpdateCache() {
	DeleteFromCache(c.ID.Hex())
	_ = c.Cache()
}

func DeleteFromCache(courseID string) {
	Cache.Del([]byte(courseID))
}

func (c *Course) AddContent(con *content.Content) {
	c.Contents = append(c.Contents, con)
}

func (c *Course) UpdateContent(con *content.Content) {
	i := c.GetContent(con.ID)
	if i >= 0 {
		c.Contents[i] = con
	}
}

func (c *Course) DeleteContent(conID primitive.ObjectID) {
	i := c.GetContent(conID)
	if i >= 0 {
		c.Contents = append(c.Contents[:i], c.Contents[i+1:]...)
	}
}

func (c *Course) AddPending(pen *pending.Pending) {
	c.Pends = append(c.Pends, pen)
}

func (c *Course) UpdatePending(pen *pending.Pending) {
	i := c.GetPending(pen.ID)
	if i >= 0 {
		c.Pends[i] = pen
	}
}

func (c *Course) DeletePending(penID primitive.ObjectID) {
	i := c.GetPending(penID)
	if i >= 0 {
		c.Pends = append(c.Pends[:i], c.Pends[i+1:]...)
	}
}

func (c *Course) AddAttachment(at *attachment.Attachment) {
	c.Inventory = append(c.Inventory, at)
}

func (c *Course) UpdateAttachment(at *attachment.Attachment) {
	i := c.GetAttachment(at.ID)
	if i >= 0 {
		c.Inventory[i] = at
	}
}

func (c *Course) DeleteAttachment(atID primitive.ObjectID) {
	i := c.GetAttachment(atID)
	if i >= 0 {
		c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
	}
}

func (c Course) IsUserProfOrTA(username string) bool {
	if c.ProfUn == username || modelUtil.ContainsInStringArray(c.TaUns, username) {
		return true
	}
	return false
}

func (c Course) IsUserProfessor(username string) bool {
	return c.ProfUn == username
}

func (c Course) IsUserStudent(username string) bool {
	return modelUtil.ContainsInStringArray(c.StdUns, username)
}

func (c Course) IsUserTA(username string) bool {
	return modelUtil.ContainsInStringArray(c.TaUns, username)
}

func (c Course) IsUserAllowedToDeleteUser(username, target string) bool {
	// professor can remove every one except his self
	if c.ProfUn == username && username != target {
		return true
	}
	// ta can remove every one except professor
	if modelUtil.ContainsInStringArray(c.TaUns, username) && c.ProfUn != target {
		return true
	}
	// every body can remove them selves except professor
	if modelUtil.ContainsInStringArray(c.StdUns, username) && username == target && c.ProfUn != username {
		return true
	}
	return false
}

func (c Course) IsUserAllowedToModifyContent(username string, content *content.Content) bool {
	// professor can remove every one except his self
	if c.ProfUn == username {
		return true
	}
	// ta can remove every one except professor
	if modelUtil.ContainsInStringArray(c.TaUns, username) && c.ProfUn != content.UploadedByUn && c.ProfUn != content.ApprovedByUn {
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
	if modelUtil.ContainsInStringArray(c.TaUns, username) && c.ProfUn != target {
		return true
	}
	// every body can remove them selves except professor
	if modelUtil.ContainsInStringArray(c.StdUns, username) && username == target {
		return true
	}
	return false
}

func (c Course) IsUserParticipateInCourse(username string) bool {
	if c.ProfUn == username || modelUtil.ContainsInStringArray(c.TaUns, username) || modelUtil.ContainsInStringArray(c.StdUns, username) {
		return true
	}
	return false
}

func (c Course) CheckCourseToken(token string) bool {
	return modelUtil.CheckTokenHash(token, c.Token)
}

func (c Course) GetContent(contentID primitive.ObjectID) int {
	for i, cnt := range c.Contents {
		if cnt.ID == contentID {
			return i
		}
	}
	return -1
}

func (c Course) GetPending(contentID primitive.ObjectID) int {
	for i, pnt := range c.Pends {
		if pnt.ID == contentID {
			return i
		}
	}
	return -1
}

func (c Course) GetAttachment(attachmentID primitive.ObjectID) int {
	for i, pnt := range c.Inventory {
		if pnt.ID == attachmentID {
			return i
		}
	}
	return -1
}

func (c *Course) RemoveComment(username string, commentID primitive.ObjectID, cnt *content.Content) (*comment.Comment, *comment.Reply, error) {
	for i, com := range cnt.Comments {
		if com.ID == commentID {
			if !c.IsUserAllowedToDeleteUserComment(username, com.AuthorUn) {
				return nil, nil, model.UserNotAllowedException{Message: "permission denied for username: " + username}
			}
			rc := cnt.Comments[i]
			cnt.Comments = append(cnt.Comments[:i], cnt.Comments[i+1:]...)
			return rc, nil, nil
		} else {
			for j, rep := range cnt.Comments[i].Replies {
				if rep.ID == commentID {
					if !c.IsUserAllowedToDeleteUserComment(username, rep.AuthorUn) {
						return nil, nil, model.UserNotAllowedException{Message: "permission denied for username: " + username}
					}
					rc := cnt.Comments[i].Replies[j]
					cnt.Comments[i].Replies = append(cnt.Comments[i].Replies[:j], cnt.Comments[i].Replies[j+1:]...)
					return nil, rc, nil
				}
			}
		}
	}
	return nil, nil, model.CommentNotFoundException{Message: "there is no comment @" + commentID.Hex()}
}

func (c *Course) IsUserAllowedToUpdateCourse(username string) error {
	if !c.IsUserProfessor(username) {
		return model.UserNotAllowedException{Message: "you can't change this course. because you are not professor"}
	}
	return nil
}

func (c *Course) IsUserAllowedToDeleteCourse(username string) error {
	if !c.IsUserProfessor(username) {
		return model.UserNotAllowedException{Message: "you are not professor"}
	}
	return nil
}

func (c *Course) IsUserAllowedToAddUserInCourse(username, token string) error {
	if c.IsUserParticipateInCourse(username) {
		return model.DuplicateUsernameException{Message: "you been been added before"}
	}
	if !c.CheckCourseToken(token) {
		return model.IncorrectTokenException{Message: "wrong course token"}
	}
	return nil
}

func (c *Course) IsUserAllowedToDeleteUserInCourse(username, targetUsername string) error {
	if !c.IsUserParticipateInCourse(targetUsername) {
		return model.UserNotFoundException{Message: "you weren't participate in course"}
	}
	if !c.IsUserAllowedToDeleteUser(username, targetUsername) {
		return model.UserNotAllowedException{Message: "you can't remove this user"}
	}
	return nil
}

func (c *Course) IsUserAllowedToPromoteUserInCourse(username, targetUsername string) error {
	if !c.IsUserProfOrTA(username) {
		return model.UserNotAllowedException{Message: "you are not professor or ta"}
	}
	if !c.IsUserStudent(targetUsername) {
		return model.UserIsNotSTDException{Message: "you are not student"}
	}
	return nil
}

func (c *Course) IsUserAllowedToDemoteUserInCourse(username, targetUsername string) error {
	if !c.IsUserProfOrTA(username) {
		return model.UserNotAllowedException{Message: "you are not professor or ta"}
	}
	if !c.IsUserTA(targetUsername) {
		return model.UserIsNotSTDException{Message: "you are not ta"}
	}
	return nil
}

func (c *Course) IsUserAllowedToInsertAttachment(username string) error {
	if !c.IsUserProfOrTA(username) {
		return model.UserNotAllowedException{Message: "you are not professor or ta"}
	}
	return nil
}

func (c *Course) IsUserAllowedToUpdateAttachment(username string) error {
	if !c.IsUserProfOrTA(username) {
		return model.UserNotAllowedException{Message: "you are not professor or ta"}
	}
	return nil
}

func (c *Course) IsUserAllowedToDeleteAttachment(username string) error {
	if !c.IsUserProfOrTA(username) {
		return model.UserNotAllowedException{Message: "you are not professor or ta"}
	}
	return nil
}

func (c *Course) IsUserAllowedToInsertContent(username string) error {
	if !c.IsUserProfOrTA(username) {
		return model.UserNotAllowedException{Message: "you are not professor or ta"}
	}
	return nil
}

func (c *Course) IsUserAllowedToUpdateContent(username string, content *content.Content) error {
	if !c.IsUserAllowedToModifyContent(username, content) {
		return model.UserNotAllowedException{Message: "you can't edit this content"}
	}
	return nil
}

func (c *Course) IsUserAllowedToDeleteContent(username string, content *content.Content) error {
	if !c.IsUserAllowedToModifyContent(username, content) {
		return model.UserNotAllowedException{Message: "you can't edit this content"}
	}
	return nil
}

func (c *Course) IsUserAllowedToInsertPending(username string) error {
	if !c.IsUserStudent(username) {
		return model.UserNotAllowedException{Message: "you are not student"}
	}
	return nil
}

func (c *Course) IsUserAllowedToUpdatePending(username string, pnd *pending.Pending) error {
	if username != pnd.UploadedByUn {
		return model.UserNotAllowedException{Message: "you can't edit this offer"}
	}
	if pnd.Status != pending.PENDING {
		return model.OfferedContentNotPendingException{Message: "this offer has been closed"}
	}
	return nil
}

func (c *Course) IsUserAllowedToDeletePending(username string, pending *pending.Pending) error {
	if !c.IsUserProfOrTA(username) && username != pending.UploadedByUn {
		return model.UserNotAllowedException{Message: "you can't remove this offer"}
	}
	return nil
}

func (c *Course) IsUserAllowedToAcceptPending(username string, ond *pending.Pending) error {
	if !c.IsUserProfOrTA(username) {
		return model.UserNotAllowedException{Message: "you can't accept this offer"}
	}
	if ond.Status != pending.PENDING {
		return model.OfferedContentNotPendingException{Message: "this offer has been closed"}
	}
	return nil
}

func (c *Course) IsUserAllowedToRejectPending(username string, pnd *pending.Pending) error {
	if !c.IsUserProfOrTA(username) {
		return model.UserNotAllowedException{Message: "you can't reject this offer"}
	}
	if pnd.Status != pending.PENDING {
		return model.OfferedContentNotPendingException{Message: "this offer has been closed"}
	}
	return nil
}

func (c *Course) IsUserAllowedToInsertComment(username string) error {
	if !c.IsUserParticipateInCourse(username) {
		return model.UserNotAllowedException{Message: "you can't write a comment on this course"}
	}
	return nil
}

func (c *Course) IsUserAllowedToUpdateComment(username, author string) error {
	if username != author {
		return model.UserNotAllowedException{Message: "you can't edit this comment on this course"}
	}
	return nil
}

func (c *Course) FilterPending(username *string, pnd *pending.Pending) *pending.Pending {
	if username == nil || (*username != pnd.UploadedByUn && !c.IsUserProfOrTA(*username)) {
		pnd.Furl = ""
		pnd.Description = ""
	}
	return pnd
}

func (c *Course) FilterPendings(username *string, pends []*pending.Pending) []*pending.Pending {
	var pnds []*pending.Pending
	for _, pn := range pends {
		pnds = append(pnds, c.FilterPending(username, pn))
	}
	return pnds
}

func (c *Course) FilterPendsOfCourse(username *string) *Course {
	c.Pends = c.FilterPendings(username, c.Pends)
	return c
}

func (c *Course) AddNewContent(authorUsername string, title string, description *string, upload []*graphql.Upload, tags []string) (*content.Content, error) {
	//check if user can insert content
	err := c.IsUserAllowedToInsertContent(authorUsername)
	if err != nil {
		return nil, err
	}

	// storing video in the Object Storage
	bucket := c.getCourseBucket()
	if err := OSD.StoreMulti(bucket, upload[0].Filename, upload); err != nil {
		return nil, err
	}
	vurl := OSD.GetURL(bucket, upload[0].Filename)

	// create a content
	cn, err := content.New(title, authorUsername, vurl, c.ID.Hex(), description, nil, tags)
	if err != nil {
		return nil, nil
	}

	// maintain consistency in cache
	c.AddContent(cn)
	c.UpdateCache()

	return cn, nil
}

func (c *Course) AddNewPending(title string, authorUsername string, upload []*graphql.Upload, description *string) (*pending.Pending, error) {

	// check if user can offer
	err := c.IsUserAllowedToInsertPending(authorUsername)
	if err != nil {
		return nil, err
	}

	// store video in Object Storage
	bucket := c.getCourseBucket()
	if err := OSD.StoreMulti(bucket, upload[0].Filename, upload); err != nil {
		return nil, err
	}
	vurl := OSD.GetURL(bucket, upload[0].Filename)

	// create a pending
	pn, err := pending.New(title, authorUsername, vurl, c.ID.Hex(), description)
	if err != nil {
		return nil, err
	}

	// insert the pending into database
	pn, err = pending.Insert(c.ID.Hex(), pn)
	if err != nil {
		return nil, err
	}
	return pn,nil
}

func (c *Course) AddNewAttachment(authorUsername string, name string, attach graphql.Upload, description *string) (*attachment.Attachment, error) {

	// check if user can insert attachment
	err := c.IsUserAllowedToInsertAttachment(authorUsername)
	if err != nil {
		return nil, err
	}

	// store attachment in object storage
	bucket := c.GetAttachmentBucket()
	if err := OSD.Store(bucket, attach.Filename, attach.File, attach.Size); err != nil {
		return nil, err
	}
	aurl := OSD.GetURL(bucket, attach.Filename)

	// create an attachment
	an, err := attachment.New(name, aurl, c.ID.Hex(), description)
	if err != nil {
		return nil, err
	}
	// insert the attachment into database
	an, err = attachment.Insert(c.ID.Hex(), an)
	if err != nil {
		return nil, err
	}
	return an,nil
}


func FilterPendsOfCourses(username *string, courses []*Course) []*Course {
	var crs []*Course
	for _, cr := range courses {
		crs = append(crs, cr.FilterPendsOfCourse(username))
	}
	return crs
}
