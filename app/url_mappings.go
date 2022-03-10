package app

import (
	"bookstore_users_api/controllers/ping"
	"bookstore_users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping) //curl -X GET localhost:8080/ping

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser) //curl -X POST localhost:8080/users -d '{"id":"123","first_name":"John"}'
}
