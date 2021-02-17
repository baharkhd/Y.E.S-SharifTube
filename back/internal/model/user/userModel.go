package user

import (
	"encoding/json"
	"github.com/coocood/freecache"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/model/course"
)

const CacheExpire = 10 * 60

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string
	Email    string
	Username string
	Password string
	Courses  []string `bson:"courses" json:"courses"`
}

var DBD DBDriver
var Cache *freecache.Cache

var DeletedAccount = &User{
	ID:       primitive.NilObjectID,
	Username: "DELETED ACCOUNT",
	Name:     "",
	Email:    "",
	Password: "",
	Courses:  nil,
}

func (u *User) enroll(courseID string) *User {
	u.Courses = append(u.Courses, courseID)
	return u
}

func (u *User) leave(CourseID string) *User {
	for i, course := range u.Courses {
		if course == CourseID {
			u.Courses = append(u.Courses[:i], u.Courses[i+1:]...)
			return u
		}
	}
	return u
}

func (u *User) updateName(name string) *User {
	u.Name = name
	return u
}

func (u *User) updateEmail(email string) *User {
	u.Email = email
	return u
}

func (u *User) updatePassword(password string) error {
	// hashing password
	hashedPass, err := hashAndSalt([]byte(password))

	if err != nil {
		return model.InternalServerException{Message: "internal server error: couldn't hash password"}
	}

	u.Password = hashedPass
	return nil
}

func GetFromCache(username string) (*User, error) {
	c, err := Cache.Get([]byte(username))
	if err == nil {
		var cr *User
		err = json.Unmarshal(c, &cr)
		if err != nil {
			return nil, model.InternalServerException{Message: err.Error()}
		}
		return cr, err
	}
	return nil, model.UserNotFoundException{Message: "user not found in cache"}
}

func DeleteFromCache(username string) {
	Cache.Del([]byte(username))
}

func (u *User) Cache() error {
	content, err := json.Marshal(u)
	if err != nil {
		return model.InternalServerException{Message: err.Error()}
	}
	err = Cache.Set([]byte(u.Username), content, CacheExpire)
	if err != nil {
		return model.InternalServerException{Message: "content couldn't cache"}
	}
	return nil
}

func (u *User) UpdateCache() {
	DeleteFromCache(u.Username)
	_ = u.Cache()
}

func DeleteUsersOfCourseFromCache(c *course.Course){
	DeleteFromCache(c.ProfUn)
	for _, ta := range c.TaUns {
		DeleteFromCache(ta)
	}
	for _, std := range c.StdUns {
		DeleteFromCache(std)
	}
}
