package user

import (
	"yes-sharifTube/graph/model"
	"yes-sharifTube/pkg/database/status"
)

func Delete(username string) error {

	if stat := UserDBD.Delete(&username); stat == status.FAILED {
		return model.InternalServerException{Message: "couldn't delete the user"}
	} else {
		return nil
	}
}

func New(name, email, username, password string) (*User, error) {

	// checking for duplicate username
	if _, stat := UserDBD.Get(&name); stat == status.SUCCESSFUL {
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
	if stat := UserDBD.Insert(user); stat == status.FAILED {
		return &User{}, model.InternalServerException{Message: "couldn't create user"}
	} else {
		return user, nil
	}
}

func Get(username string) (*User, error) {
	if target, stat := UserDBD.Get(&username); stat == status.FAILED {
		return nil, model.UserNotFoundException{Message: "couldn't find the requested user"}
	} else {
		return target, nil
	}
}
