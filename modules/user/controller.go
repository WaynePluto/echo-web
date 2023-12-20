package user

import (
	"echo-web/custom"

	"github.com/labstack/echo/v4"
)

// addUser godoc
//
//	@Summary			addUser
//	@Description
//	@Tags					user
//	@Accept				json
//	@Produce			json
//	@Param				userinfo	body	UserModel	true	"user info object"
//	@Success			200 {object} UserModel
//	@Failure			400 {object} interface{}
//	@Router				/users [post]
func addUser(c echo.Context) error {
	cc := c.(*custom.Context)
	var user UserModel
	err := cc.BindDTO(&user)
	if err != nil {
		return cc.ClientErr(err)
	}
	cc.Logger().Infof("user dto: %v", user)
	// service

	return cc.Send(user)
}

// getUserById godoc
//
//	@Summary		getUserById
//	@Description
//	@Tags				user
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"user ID"
//	@Success		200 {object} UserModel
//	@Failure		400 {object} interface{}
//	@Router			/users/{id} [get]
func getUserById(c echo.Context) error {
	cc := c.(*custom.Context)
	var idDTO IdDTO
	cc.BindDTO(&idDTO)
	cc.Logger().Infof("user id: %v", idDTO)
	return cc.Send(UserModel{Num: idDTO.ID})
}

// updateUserById godoc
//
//	@Summary		updateUserById
//	@Description
//	@Tags				user
//	@Accept			json
//	@Produce		json
//	@Param			id	path	number	true	"user ID"
//	@Success		200 {object} UserModel
//	@Failure		400 {object} interface{}
//	@Router			/users/{id} [put]
func updateUserById(c echo.Context) error {
	cc := c.(*custom.Context)
	return cc.Send(UserModel{})
}

// updateUserById godoc
//
//	@Summary		updateUserById
//	@Description
//	@Tags				user
//	@Accept			json
//	@Produce		json
//	@Param			id	path	number	true	"user ID"
//	@Success		200 {object} UserModel
//	@Failure		400 {object} interface{}
//	@Router			/users/{id} [delete]
func deleteUserById(c echo.Context) error {
	cc := c.(*custom.Context)
	return cc.Send(UserModel{})
}

// getUserList godoc
//
//	@Summary			getUserList
//	@Description
//	@Tags					user
//	@Accept				json
//	@Produce			json
//	@Param				page	query	PageDTO	true	"page"
//	@Success			200 {object} []UserModel
//	@Failure			400 {object} interface{}
//	@Router				/users [get]
func getUserList(c echo.Context) error {
	cc := c.(*custom.Context)

	var pageDTO PageDTO
	err := cc.BindDTO(&pageDTO)
	if err != nil {
		return cc.ClientErr(err)
	}

	return cc.Send([]UserModel{{Num: pageDTO.Page}})
}
