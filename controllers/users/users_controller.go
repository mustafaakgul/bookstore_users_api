package users

import (
	"bookstore_users_api/domain/users"
	"bookstore_users_api/services"
	"bookstore_users_api/utils/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	counter int
)

func CreateUser(c *gin.Context) {
	var user users.User

	/*bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//TODO: handle error
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		//TODO: handle JSON error
		fmt.Println(err.Error())
		return
	} shoulldbind json usstekiler yerine kullanlablr*/

	//curl -X POST localhost:8080/users -d  '{"id": 123, "first_name":"John", "last_name":"Doe", "email":"asda@gmail.com"}'
	if err := c.ShouldBindJSON(&user); err != nil {
		/*restErr := errors.RestErr{
			Message: "Invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}*/
		restErr := errors.NewBadRequestError("Invalid JSON Body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(user)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user_id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
