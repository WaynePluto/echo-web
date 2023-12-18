package user

import (
	"echo-web/custom"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" validate:"required" example:"zhangsan"`
	Phone string `json:"phone" validate:"required" example:"13600000000"`
	Pass  string `json:"pass"`
}

func Mount(e *echo.Echo) {
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
}

// SaveUser godoc
//
//	@Summary			addUser
//	@Description	save user
//	@Tags					user
//	@Accept				json
//	@Produce			json
//	@Param				userinfo	body	User	true	"user info object"
//	@Success			200 {object} User
//	@Failure			400 {object} interface{}
//	@Router				/users [post]
func saveUser(c echo.Context) error {
	cc := c.(*custom.Context)
	var user User

	err := cc.SetStruct(&user)
	if err != nil {
		return cc.ClientErr(err)
	}

	return cc.Send(user)
}

// GetUser godoc
//
//	@Summary		getUserById
//	@Description	get user info
//	@Tags				user
//	@Accept			json
//	@Produce		json
//	@Param			id	path	number	true	"user ID"
//	@Success		200 {object} User
//	@Failure		400 {object} interface{}
//	@Router			/users/:id [get]
func getUser(c echo.Context) error {
	cc := c.(*custom.Context)
	// id := cc.Param("id")
	return cc.Send(&User{})
}
