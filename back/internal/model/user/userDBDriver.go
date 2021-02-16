package user

import (
	"yes-sharifTube/pkg/database/status"
)

type DBDriver interface {
	Delete(name *string) status.QueryStatus
	Insert(user *User) status.QueryStatus
	InsertExact(user *User) status.QueryStatus
	Get(name *string) (*User, status.QueryStatus)
	Update(target string, user *User) status.QueryStatus
	GetAll(start, amount int64) ([]*User, status.QueryStatus)
	Replace(target *string,toBe *User) status.QueryStatus
    Enroll(username, courseID string) status.QueryStatus
    Leave(username, courseID string) status.QueryStatus
}
