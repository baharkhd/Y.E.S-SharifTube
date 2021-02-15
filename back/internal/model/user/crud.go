package user

import (
	"yes-sharifTube/graph/model"
	"yes-sharifTube/internal/model/course"
	"yes-sharifTube/pkg/database/status"
)

/*	the actual implementation of CRUD for user model is here
	we also added getAll method due to our certain needs
*/
/*	we use status.QueryStatus as a statusCode for our controllers
	we use status.FAILED to return a failed status and
	status.SUCCESSFUL to return a successful status (obviously)
*/

func GetAll(start, amount int64) ([]*User, error) {
	all, err := DBD.GetAll(start, amount)
	if err == status.FAILED {
		return nil, model.InternalServerException{Message: "couldn't fetch the required users"}
	}
	return all, nil
}

func Update(targetUsername string, toBe model.EditedUser) (*User, error) {
	targetUser := newFrom(toBe)
	return update(targetUsername, targetUser)
}

func update(targetUsername string, targetUser User) (*User, error) {
	// updating the database
	if stat := DBD.Update(targetUsername, &targetUser); stat == status.FAILED {

		// checking if the target user exists
		_, stat2 := DBD.Get(&targetUsername)
		if stat2 == status.FAILED {
			return nil, model.UserNotFoundException{Message: "target Doesnt exist"}
		}
		// no clue why query failed
		return nil, model.InternalServerException{Message: "couldn't update the user"}
	} else {

		// update user in cache if exists
		DeleteFromCache(targetUser.Username)
		_ = targetUser.Cache()

		return &targetUser, nil
	}
}

func newFrom(toBe model.EditedUser) User {
	// filling the update fields of the user
	var targetUser = User{}
	if toBe.Name != nil {
		targetUser.updateName(*toBe.Name)
	}
	if toBe.Password != nil {
		_ = targetUser.updatePassword(*(toBe.Password))
	}
	if toBe.Email != nil {
		targetUser.updateEmail(*toBe.Email)
	}
	return targetUser
}

func Delete(username string) error {

	// delete user from its courses
	usr, err := Get(username)
	if err != nil {
		return model.UserNotFoundException{Message: "user couldn't found"}
	}
	for _, cID := range usr.Courses {
		c, err := course.Get(cID)
		if err == nil {
			if c.IsUserProfessor(username) {
				err = course.Delete(c)
				if err != nil {
					return model.InternalServerException{Message: "course of user couldn't delete"}
				}
			} else {
				_, err = course.DeleteUser(username, c)
				if err != nil {
					return model.InternalServerException{Message: "user couldn't delete in the course"}
				}
			}
		}
	}
	// delete user from database
	if stat := DBD.Delete(&username); stat == status.FAILED {
		return model.InternalServerException{Message: "couldn't delete the user"}
	} else {
		// delete from cache if exists
		DeleteFromCache(username)
		return nil
	}
}

func New(name, email, username, password string) (*User, error) {

	// checking for duplicate username
	if _, stat := DBD.Get(&username); stat == status.SUCCESSFUL {
		return nil, model.DuplicateUsernameException{}
	}

	// hashing password
	hashedPass, err := hashAndSalt([]byte(password))
	if err != nil {
		return nil, model.InternalServerException{Message: "internal server error: couldn't hash password"}
	}

	user := &User{
		Name:     name,
		Email:    email,
		Username: username,
		Password: hashedPass,
		Courses:  []string{},
	}

	// inserting into the database
	if stat := DBD.Insert(user); stat == status.FAILED {
		return &User{}, model.InternalServerException{Message: "couldn't create user"}
	} else {
		return user, nil
	}
}

func Get(username string) (*User, error) {

	// checking to be in cache first
	u, err := GetFromCache(username)
	if err == nil {
		return u, nil
	}

	// if not exists, get from database
	if target, stat := DBD.Get(&username); stat == status.FAILED {
		return nil, model.UserNotFoundException{Message: "couldn't find the requested user"}
	} else {

		// add the content to cache
		_ = target.Cache()

		return target, nil
	}
}

func SetDeletedAccount() error {
	// checking for duplicate username
	if _, stat := DBD.Get(&DeletedAccount.Username); stat == status.FAILED {
		if stat = DBD.InsertExact(DeletedAccount); stat == status.FAILED {
			return model.InternalServerException{Message: "couldn't create deleted account user"}
		}
	}
	return nil
}

func GetA(usernames []string) []*User {
	var users []*User
	for i, _ := range usernames {
		users = append(users, GetS(usernames[i]))
	}
	return users
}

func GetS(username string) *User {
	if username == "" {
		return nil
	}
	u, err := Get(username)
	if err != nil {
		u = DeletedAccount
	}
	return u
}

func (u *User) Enroll(courseID string) *User {
	u.enroll(courseID)
	return u
}

func (u *User) Leave(CourseID string) *User {
	u.leave(CourseID)
	return u
}

func (u *User) UpdateName(name string) (*User, error) {
	u.updateName(name)
	return update(u.Username, User{Name: u.Name})
}

func (u *User) UpdateEmail(email string) (*User, error) {
	u.updateEmail(email)
	return update(u.Username, User{Email: u.Email})
}

func (u *User) UpdatePassword(password string) (*User, error) {
	if err := u.updatePassword(password); err != nil {
		return nil, err
	}
	return update(u.Username, User{Password: u.Password})
}
