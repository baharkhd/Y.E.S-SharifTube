package database

const UserNotFoundMessage = "there is no User with username: "
const CourseNotFoundMessage = "there is no course @"
const ContentNotFoundMessage = "there is no content @"
const AttachmentNotFoundMessage = "there is no attachment @"
const PendingNotFoundMessage = "there is no pending @"
const CommentNotFoundMessage = "there is no comment @"
const UserNotAllowedMessage = "permission denied for username: "
const FieldsEmptyMessage = "cannot set empty fields"
const InternalMessage = "fatal error in database /n"
const DuplicateUsernameMessage = "received username is duplicated"
const IncorrectTokenMessage = "incorrect course token"
const UserIsNotTAMessage = "there is no TA in this course with username: "
const UserIsNotSTDMessage = "there is no student in this course with username: "
const OfferedContentRejectedMessage = "there is no accepted offer @"
const OfferedContentNotPendingMessage = "there is no pending offer @"

type InternalDBError struct {
	Message string
}

func (e *InternalDBError) Error() string {
	return e.Message
}

func ThrowInternalDBException(mongoErr string) error {
	return &InternalDBError{Message: InternalMessage + mongoErr}
}

type AllFieldsEmpty struct {
	Message string
}

func (e *AllFieldsEmpty) Error() string {
	return e.Message
}

func ThrowAllFieldsEmptyException() error {
	return &AllFieldsEmpty{Message: FieldsEmptyMessage}
}

type DuplicateUsername struct {
	Message string
}

func (e *DuplicateUsername) Error() string {
	return e.Message
}

func ThrowDuplicateUsernameException() error {
	return &AllFieldsEmpty{Message: DuplicateUsernameMessage}
}

type UserNotFound struct {
	Message string
}

func (e *UserNotFound) Error() string {
	return e.Message
}

func ThrowUserNotFoundException(username string) error {
	return &UserNotFound{Message: UserNotFoundMessage + username}
}

type UserNotAllowed struct {
	Message string
}

func (e *UserNotAllowed) Error() string {
	return e.Message
}

func ThrowUserNotAllowedException(username string) error {
	return &UserNotAllowed{Message: UserNotAllowedMessage + username}
}

type CourseNotFound struct {
	Message string
}

func (e *CourseNotFound) Error() string {
	return e.Message
}

func ThrowCourseNotFoundException(courseID string) error {
	return &CourseNotFound{Message: CourseNotFoundMessage + courseID}
}

type IncorrectToken struct {
	Message string
}

func (e *IncorrectToken) Error() string {
	return e.Message
}

func ThrowIncorrectTokenException() error {
	return &IncorrectToken{Message: IncorrectTokenMessage}
}

type UserIsNotTA struct {
	Message string
}

func (e *UserIsNotTA) Error() string {
	return e.Message
}

func ThrowUserIsNotTAException(username string) error {
	return &UserIsNotTA{Message: UserIsNotTAMessage + username}
}

type UserIsNotSTD struct {
	Message string
}

func (e *UserIsNotSTD) Error() string {
	return e.Message
}

func ThrowUserIsNotSTDException(username string) error {
	return &UserIsNotSTD{Message: UserIsNotSTDMessage + username}
}

type ContentNotFound struct {
	Message string
}

func (e *ContentNotFound) Error() string {
	return e.Message
}

func ThrowContentNotFoundException(contentID string) error {
	return &ContentNotFound{Message: ContentNotFoundMessage + contentID}
}

type AttachmentNotFound struct {
	Message string
}

func (e *AttachmentNotFound) Error() string {
	return e.Message
}

func ThrowAttachmentNotFoundException(attachmentID string) error {
	return &AttachmentNotFound{Message: AttachmentNotFoundMessage + attachmentID}
}

type PendingNotFound struct {
	Message string
}

func (e *PendingNotFound) Error() string {
	return e.Message
}

func ThrowPendingNotFoundException(pendingID string) error {
	return &PendingNotFound{Message: PendingNotFoundMessage + pendingID}
}

type CommentNotFound struct {
	Message string
}

func (e *CommentNotFound) Error() string {
	return e.Message
}

func ThrowCommentNotFoundException(commentID string) error {
	return &CommentNotFound{Message: CommentNotFoundMessage + commentID}
}

type OfferedContentNotPending struct {
	Message string
}

func (e *OfferedContentNotPending) Error() string {
	return e.Message
}

func ThrowOfferedContentNotPendingException(pendingID string) error {
	return &OfferedContentNotPending{Message: OfferedContentRejectedMessage + pendingID}
}
