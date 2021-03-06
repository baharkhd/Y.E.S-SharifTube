package model

const EmptyKeyErrorMessage =  "all of your fields are empty"

func (e InternalServerException) Error() string {
	return e.Message
}
func (e EmptyFieldsException) Error() string {
	return e.Message
}
func (e RegexMismatchException) Error() string {
	return e.Message
}
func (e DuplicateUsernameException) Error() string {
	return e.Message
}
func (e UserNotFoundException) Error() string {
	return e.Message
}
func (e UserNotAllowedException) Error() string {
	return e.Message
}
func (e UserPassMissMatchException) Error() string {
	return e.Message
}
func (e CourseNotFoundException) Error() string {
	return e.Message
}
func (e IncorrectTokenException) Error() string {
	return e.Message
}
func (e UserIsNotTAException) Error() string {
	return e.Message
}
func (e UserIsNotSTDException) Error() string {
	return e.Message
}
func (e ContentNotFoundException) Error() string {
	return e.Message
}
func (e AttachmentNotFoundException) Error() string {
	return e.Message
}
func (e PendingNotFoundException) Error() string {
	return e.Message
}
func (e OfferedContentNotPendingException) Error() string {
	return e.Message
}
func (e CommentNotFoundException) Error() string {
	return e.Message
}

func (e FileAlreadyExistsException) Error() string {
	return e.Message
}
