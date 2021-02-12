package user

import (
	"yes-sharifTube/graph/model"
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

func Update(targetUsername string, toBe model.EditedUser) (*User,error) {

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
		targetUser.updateEmail(*toBe.Name)
	}
	return targetUser
}


func Delete(username string) error {

	if stat := DBD.Delete(&username); stat == status.FAILED {
		return model.InternalServerException{Message: "couldn't delete the user"}
	} else {
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
	if target, stat := DBD.Get(&username); stat == status.FAILED {
		return nil, model.UserNotFoundException{Message: "couldn't find the requested user"}
	} else {
		return target, nil
	}
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
	return update(u.Username,User{Name: u.Name})
}
func (u *User) UpdateEmail(email string) (*User, error) {
	u.updateEmail(email)
	return update(u.Username, User{Email: u.Email})
}

func (u *User) UpdatePassword(password string) (*User, error) {
	if err := u.updatePassword(password); err!=nil{
		return nil, err
	}
	return update(u.Username,User{Password: u.Password})
}

