package services

import (
	"bookstore_users_api/domain/users"
	"bookstore_users_api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	/*user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return nil, errors.NewBadRequestError("invalid email address")
	}*/
	//users.Validate(&user) goes Validate function, user.Validate() goes Validate method
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil //returns pointer of the user and nil error
}
