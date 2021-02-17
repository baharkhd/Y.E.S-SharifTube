// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

type AddUserToCoursePayload interface {
	IsAddUserToCoursePayload()
}

type CreateCommentPayLoad interface {
	IsCreateCommentPayLoad()
}

type CreateCoursePayload interface {
	IsCreateCoursePayload()
}

type CreateUserPayload interface {
	IsCreateUserPayload()
}

type DeleteAttachmentPayLoad interface {
	IsDeleteAttachmentPayLoad()
}

type DeleteCommentPayLoad interface {
	IsDeleteCommentPayLoad()
}

type DeleteContentPayLoad interface {
	IsDeleteContentPayLoad()
}

type DeleteCoursePayload interface {
	IsDeleteCoursePayload()
}

type DeleteOfferedContentPayLoad interface {
	IsDeleteOfferedContentPayLoad()
}

type DeleteUserFromCoursePayload interface {
	IsDeleteUserFromCoursePayload()
}

type DeleteUserPayload interface {
	IsDeleteUserPayload()
}

type DemoteToSTDPayload interface {
	IsDemoteToSTDPayload()
}

type EditAttachmentPayLoad interface {
	IsEditAttachmentPayLoad()
}

type EditCommentPayLoad interface {
	IsEditCommentPayLoad()
}

type EditContentPayLoad interface {
	IsEditContentPayLoad()
}

type EditOfferedContentPayLoad interface {
	IsEditOfferedContentPayLoad()
}

type Exception interface {
	IsException()
}

type LoginPayload interface {
	IsLoginPayload()
}

type OfferContentPayLoad interface {
	IsOfferContentPayLoad()
}

type PromoteToTAPayload interface {
	IsPromoteToTAPayload()
}

type UpdateCourseInfoPayload interface {
	IsUpdateCourseInfoPayload()
}

type UpdateUserPayload interface {
	IsUpdateUserPayload()
}

type UploadAttachmentPayLoad interface {
	IsUploadAttachmentPayLoad()
}

type UploadContentPayLoad interface {
	IsUploadContentPayLoad()
}

type AcceptedPending struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Tags        []string `json:"tags"`
	Message     *string  `json:"message"`
}

type Attachment struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Aurl        string  `json:"aurl"`
	Description *string `json:"description"`
	Timestamp   int     `json:"timestamp"`
	CourseID    string  `json:"courseID"`
}

func (Attachment) IsUploadAttachmentPayLoad() {}
func (Attachment) IsEditAttachmentPayLoad()   {}
func (Attachment) IsDeleteAttachmentPayLoad() {}

type AttachmentNotFoundException struct {
	Message string `json:"message"`
}

func (AttachmentNotFoundException) IsException()               {}
func (AttachmentNotFoundException) IsEditAttachmentPayLoad()   {}
func (AttachmentNotFoundException) IsDeleteAttachmentPayLoad() {}

type Comment struct {
	ID        string   `json:"id"`
	Author    *User    `json:"author"`
	Body      string   `json:"body"`
	Timestamp int      `json:"timestamp"`
	Replies   []*Reply `json:"replies"`
	ContentID string   `json:"contentID"`
}

func (Comment) IsCreateCommentPayLoad() {}
func (Comment) IsEditCommentPayLoad()   {}
func (Comment) IsDeleteCommentPayLoad() {}

type CommentNotFoundException struct {
	Message string `json:"message"`
}

func (CommentNotFoundException) IsException()            {}
func (CommentNotFoundException) IsCreateCommentPayLoad() {}
func (CommentNotFoundException) IsEditCommentPayLoad()   {}
func (CommentNotFoundException) IsDeleteCommentPayLoad() {}

type Content struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Timestamp   int        `json:"timestamp"`
	UploadedBy  *User      `json:"uploadedBY"`
	ApprovedBy  *User      `json:"approvedBY"`
	Vurl        string     `json:"vurl"`
	Comments    []*Comment `json:"comments"`
	Tags        []string   `json:"tags"`
	CourseID    string     `json:"courseID"`
}

func (Content) IsUploadContentPayLoad() {}
func (Content) IsEditContentPayLoad()   {}
func (Content) IsDeleteContentPayLoad() {}

type ContentNotFoundException struct {
	Message string `json:"message"`
}

func (ContentNotFoundException) IsException()            {}
func (ContentNotFoundException) IsEditContentPayLoad()   {}
func (ContentNotFoundException) IsDeleteContentPayLoad() {}
func (ContentNotFoundException) IsCreateCommentPayLoad() {}
func (ContentNotFoundException) IsEditCommentPayLoad()   {}
func (ContentNotFoundException) IsDeleteCommentPayLoad() {}

type Course struct {
	ID        string        `json:"id"`
	Title     string        `json:"title"`
	Summary   *string       `json:"summary"`
	CreatedAt int           `json:"createdAt"`
	Token     string        `json:"token"`
	Prof      *User         `json:"prof"`
	Tas       []*User       `json:"tas"`
	Pends     []*Pending    `json:"pends"`
	Students  []*User       `json:"students"`
	Contents  []*Content    `json:"contents"`
	Inventory []*Attachment `json:"inventory"`
}

func (Course) IsCreateCoursePayload()         {}
func (Course) IsUpdateCourseInfoPayload()     {}
func (Course) IsDeleteCoursePayload()         {}
func (Course) IsAddUserToCoursePayload()      {}
func (Course) IsDeleteUserFromCoursePayload() {}
func (Course) IsPromoteToTAPayload()          {}
func (Course) IsDemoteToSTDPayload()          {}

type CourseNotFoundException struct {
	Message string `json:"message"`
}

func (CourseNotFoundException) IsException()                   {}
func (CourseNotFoundException) IsUpdateCourseInfoPayload()     {}
func (CourseNotFoundException) IsDeleteCoursePayload()         {}
func (CourseNotFoundException) IsAddUserToCoursePayload()      {}
func (CourseNotFoundException) IsDeleteUserFromCoursePayload() {}
func (CourseNotFoundException) IsPromoteToTAPayload()          {}
func (CourseNotFoundException) IsDemoteToSTDPayload()          {}
func (CourseNotFoundException) IsUploadContentPayLoad()        {}
func (CourseNotFoundException) IsEditContentPayLoad()          {}
func (CourseNotFoundException) IsDeleteContentPayLoad()        {}
func (CourseNotFoundException) IsUploadAttachmentPayLoad()     {}
func (CourseNotFoundException) IsEditAttachmentPayLoad()       {}
func (CourseNotFoundException) IsDeleteAttachmentPayLoad()     {}
func (CourseNotFoundException) IsOfferContentPayLoad()         {}
func (CourseNotFoundException) IsEditOfferedContentPayLoad()   {}
func (CourseNotFoundException) IsDeleteOfferedContentPayLoad() {}

type DuplicateUsernameException struct {
	Message string `json:"message"`
}

func (DuplicateUsernameException) IsException()              {}
func (DuplicateUsernameException) IsCreateUserPayload()      {}
func (DuplicateUsernameException) IsAddUserToCoursePayload() {}

type EditAttachment struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type EditContent struct {
	Title       *string  `json:"title"`
	Description *string  `json:"description"`
	Tags        []string `json:"tags"`
}

type EditedComment struct {
	Body *string `json:"body"`
}

type EditedCourse struct {
	Title   *string `json:"title"`
	Summary *string `json:"summary"`
	Token   *string `json:"token"`
}

type EditedPending struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type EditedUser struct {
	Password *string `json:"password"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
}

type EmptyFieldsException struct {
	Message string `json:"message"`
}

func (EmptyFieldsException) IsException()                 {}
func (EmptyFieldsException) IsUpdateCourseInfoPayload()   {}
func (EmptyFieldsException) IsEditContentPayLoad()        {}
func (EmptyFieldsException) IsEditAttachmentPayLoad()     {}
func (EmptyFieldsException) IsEditOfferedContentPayLoad() {}
func (EmptyFieldsException) IsEditCommentPayLoad()        {}

type FileAlreadyExistsException struct {
	Message string `json:"message"`
}

func (FileAlreadyExistsException) IsException() {}

type IncorrectTokenException struct {
	Message string `json:"message"`
}

func (IncorrectTokenException) IsException()              {}
func (IncorrectTokenException) IsAddUserToCoursePayload() {}

type InternalServerException struct {
	Message string `json:"message"`
}

func (InternalServerException) IsException()                   {}
func (InternalServerException) IsCreateUserPayload()           {}
func (InternalServerException) IsUpdateUserPayload()           {}
func (InternalServerException) IsDeleteUserPayload()           {}
func (InternalServerException) IsLoginPayload()                {}
func (InternalServerException) IsCreateCoursePayload()         {}
func (InternalServerException) IsUpdateCourseInfoPayload()     {}
func (InternalServerException) IsDeleteCoursePayload()         {}
func (InternalServerException) IsAddUserToCoursePayload()      {}
func (InternalServerException) IsDeleteUserFromCoursePayload() {}
func (InternalServerException) IsPromoteToTAPayload()          {}
func (InternalServerException) IsDemoteToSTDPayload()          {}
func (InternalServerException) IsUploadContentPayLoad()        {}
func (InternalServerException) IsEditContentPayLoad()          {}
func (InternalServerException) IsDeleteContentPayLoad()        {}
func (InternalServerException) IsUploadAttachmentPayLoad()     {}
func (InternalServerException) IsEditAttachmentPayLoad()       {}
func (InternalServerException) IsDeleteAttachmentPayLoad()     {}
func (InternalServerException) IsOfferContentPayLoad()         {}
func (InternalServerException) IsEditOfferedContentPayLoad()   {}
func (InternalServerException) IsDeleteOfferedContentPayLoad() {}
func (InternalServerException) IsCreateCommentPayLoad()        {}
func (InternalServerException) IsEditCommentPayLoad()          {}
func (InternalServerException) IsDeleteCommentPayLoad()        {}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type OfferedContentNotPendingException struct {
	Message string `json:"message"`
}

func (OfferedContentNotPendingException) IsException()                   {}
func (OfferedContentNotPendingException) IsEditOfferedContentPayLoad()   {}
func (OfferedContentNotPendingException) IsDeleteOfferedContentPayLoad() {}

type OperationSuccessfull struct {
	Message string `json:"message"`
}

func (OperationSuccessfull) IsDeleteUserPayload() {}

type Pending struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Status      Status  `json:"status"`
	Timestamp   int     `json:"timestamp"`
	UploadedBy  *User   `json:"uploadedBY"`
	Furl        string  `json:"furl"`
	CourseID    string  `json:"courseID"`
	Message     *string `json:"message"`
}

func (Pending) IsOfferContentPayLoad()         {}
func (Pending) IsEditOfferedContentPayLoad()   {}
func (Pending) IsDeleteOfferedContentPayLoad() {}

type PendingFilter struct {
	CourseID         *string `json:"courseID"`
	Status           *Status `json:"status"`
	UploaderUsername *string `json:"uploaderUsername"`
}

type PendingNotFoundException struct {
	Message string `json:"message"`
}

func (PendingNotFoundException) IsException()                   {}
func (PendingNotFoundException) IsEditOfferedContentPayLoad()   {}
func (PendingNotFoundException) IsDeleteOfferedContentPayLoad() {}

type RegexMismatchException struct {
	Message string `json:"message"`
}

func (RegexMismatchException) IsException()                 {}
func (RegexMismatchException) IsCreateUserPayload()         {}
func (RegexMismatchException) IsCreateCoursePayload()       {}
func (RegexMismatchException) IsUpdateCourseInfoPayload()   {}
func (RegexMismatchException) IsUploadContentPayLoad()      {}
func (RegexMismatchException) IsEditContentPayLoad()        {}
func (RegexMismatchException) IsUploadAttachmentPayLoad()   {}
func (RegexMismatchException) IsEditAttachmentPayLoad()     {}
func (RegexMismatchException) IsOfferContentPayLoad()       {}
func (RegexMismatchException) IsEditOfferedContentPayLoad() {}
func (RegexMismatchException) IsCreateCommentPayLoad()      {}
func (RegexMismatchException) IsEditCommentPayLoad()        {}

type RejectedPending struct {
	Message *string `json:"message"`
}

type Reply struct {
	ID        string `json:"id"`
	Author    *User  `json:"author"`
	Body      string `json:"body"`
	Timestamp int    `json:"timestamp"`
	CommentID string `json:"commentID"`
}

func (Reply) IsCreateCommentPayLoad() {}
func (Reply) IsEditCommentPayLoad()   {}
func (Reply) IsDeleteCommentPayLoad() {}

type TargetAttachment struct {
	Name        string         `json:"name"`
	Attach      graphql.Upload `json:"attach"`
	Description *string        `json:"description"`
}

type TargetComment struct {
	Body string `json:"body"`
}

type TargetContent struct {
	Title       string            `json:"title"`
	Description *string           `json:"description"`
	Video       []*graphql.Upload `json:"video"`
	Tags        []string          `json:"tags"`
}

type TargetCourse struct {
	Title   string  `json:"title"`
	Summary *string `json:"summary"`
	Token   string  `json:"token"`
}

type TargetPending struct {
	Title       string            `json:"title"`
	Description *string           `json:"description"`
	Video       []*graphql.Upload `json:"video"`
}

type TargetUser struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Name     *string `json:"name"`
	Email    *string `json:"email"`
}

type Token struct {
	Token string `json:"token"`
}

func (Token) IsLoginPayload() {}

type User struct {
	ID        string   `json:"id"`
	Username  string   `json:"username"`
	Name      *string  `json:"name"`
	Email     *string  `json:"email"`
	CourseIDs []string `json:"courseIDs"`
}

func (User) IsCreateUserPayload() {}
func (User) IsUpdateUserPayload() {}
func (User) IsDeleteUserPayload() {}

type UserIsNotSTDException struct {
	Message string `json:"message"`
}

func (UserIsNotSTDException) IsException()          {}
func (UserIsNotSTDException) IsPromoteToTAPayload() {}

type UserIsNotTAException struct {
	Message string `json:"message"`
}

func (UserIsNotTAException) IsException()          {}
func (UserIsNotTAException) IsDemoteToSTDPayload() {}

type UserNotAllowedException struct {
	Message string `json:"message"`
}

func (UserNotAllowedException) IsException()                   {}
func (UserNotAllowedException) IsUpdateUserPayload()           {}
func (UserNotAllowedException) IsDeleteUserPayload()           {}
func (UserNotAllowedException) IsUpdateCourseInfoPayload()     {}
func (UserNotAllowedException) IsDeleteCoursePayload()         {}
func (UserNotAllowedException) IsAddUserToCoursePayload()      {}
func (UserNotAllowedException) IsDeleteUserFromCoursePayload() {}
func (UserNotAllowedException) IsPromoteToTAPayload()          {}
func (UserNotAllowedException) IsDemoteToSTDPayload()          {}
func (UserNotAllowedException) IsUploadContentPayLoad()        {}
func (UserNotAllowedException) IsEditContentPayLoad()          {}
func (UserNotAllowedException) IsDeleteContentPayLoad()        {}
func (UserNotAllowedException) IsUploadAttachmentPayLoad()     {}
func (UserNotAllowedException) IsEditAttachmentPayLoad()       {}
func (UserNotAllowedException) IsDeleteAttachmentPayLoad()     {}
func (UserNotAllowedException) IsOfferContentPayLoad()         {}
func (UserNotAllowedException) IsEditOfferedContentPayLoad()   {}
func (UserNotAllowedException) IsDeleteOfferedContentPayLoad() {}
func (UserNotAllowedException) IsCreateCommentPayLoad()        {}
func (UserNotAllowedException) IsEditCommentPayLoad()          {}
func (UserNotAllowedException) IsDeleteCommentPayLoad()        {}

type UserNotFoundException struct {
	Message string `json:"message"`
}

func (UserNotFoundException) IsException()                   {}
func (UserNotFoundException) IsUpdateUserPayload()           {}
func (UserNotFoundException) IsDeleteUserPayload()           {}
func (UserNotFoundException) IsCreateCoursePayload()         {}
func (UserNotFoundException) IsUpdateCourseInfoPayload()     {}
func (UserNotFoundException) IsDeleteCoursePayload()         {}
func (UserNotFoundException) IsAddUserToCoursePayload()      {}
func (UserNotFoundException) IsDeleteUserFromCoursePayload() {}
func (UserNotFoundException) IsPromoteToTAPayload()          {}
func (UserNotFoundException) IsDemoteToSTDPayload()          {}
func (UserNotFoundException) IsUploadContentPayLoad()        {}
func (UserNotFoundException) IsEditContentPayLoad()          {}
func (UserNotFoundException) IsDeleteContentPayLoad()        {}
func (UserNotFoundException) IsUploadAttachmentPayLoad()     {}
func (UserNotFoundException) IsEditAttachmentPayLoad()       {}
func (UserNotFoundException) IsDeleteAttachmentPayLoad()     {}
func (UserNotFoundException) IsOfferContentPayLoad()         {}
func (UserNotFoundException) IsEditOfferedContentPayLoad()   {}
func (UserNotFoundException) IsDeleteOfferedContentPayLoad() {}
func (UserNotFoundException) IsCreateCommentPayLoad()        {}
func (UserNotFoundException) IsEditCommentPayLoad()          {}
func (UserNotFoundException) IsDeleteCommentPayLoad()        {}

type UserPassMissMatchException struct {
	Message string `json:"message"`
}

func (UserPassMissMatchException) IsException()    {}
func (UserPassMissMatchException) IsLoginPayload() {}

type Status string

const (
	StatusPending  Status = "PENDING"
	StatusAccepted Status = "ACCEPTED"
	StatusRejected Status = "REJECTED"
)

var AllStatus = []Status{
	StatusPending,
	StatusAccepted,
	StatusRejected,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusPending, StatusAccepted, StatusRejected:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
