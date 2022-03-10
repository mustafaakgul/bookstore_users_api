package users

import (
	"bookstore_users_api/utils/date_utils"
	"bookstore_users_api/utils/errors"
	"fmt"
)

/*func Get(userId int64) (*User, *errors.RestErr) {
	return nil, nil
}*/

var (
	usersDB = make(map[int64]*User)
)

// We user pointer as we want to modify the existing user directly
func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	//now := time.Now() //	now := time.Now().UTC()
	//user.DateCreated = now.Format("2006-01-02T15:04:05Z")
	user.DateCreated = date_utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
