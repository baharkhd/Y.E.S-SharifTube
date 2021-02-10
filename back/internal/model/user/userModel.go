package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/pkg/database/mongodb"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string
	Email    string
	Username string
	Password string
	Courses  []string
}

var UserDBD *mongodb.UserMongoDriver


func (u *User) Enroll(courseID string) *User {
	u.Courses = append(u.Courses, courseID)
	return u
}

func (u *User) Leave(CourseID string) *User {
	for i, course := range u.Courses {
		if course == CourseID {
			u.Courses = append(u.Courses[:i], u.Courses[i+1:]...)
			return u
		}
	}
	return u
}

func (u *User) UpdateName(name string) *User {
	u.Name = name
	return u
}
func (u *User) UpdateEmail(email string) *User {
	u.Email=email
	return u
}

func (u *User) UpdatePassword(password string) error {
	// hashing password
	hashedPass, err := hashAndSalt([]byte(password))

	if err != nil {
		return model.InternalServerException{Message: "internal server error: couldn't hash password"}
	}

	u.Password = hashedPass
	return nil
}