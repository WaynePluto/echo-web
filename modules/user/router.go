package user

import "github.com/labstack/echo/v4"

func SetRouter(e *echo.Echo) {
	e.POST("/users", addUser)              // 创建
	e.GET("/users/:id", getUserById)       // 获取单个
	e.PUT("/users/:id", updateUserById)    // 修改
	e.PATCH("/users/:id", updateUserById)  // 部分修改
	e.DELETE("/users/:id", deleteUserById) // 删除
	e.GET("/users", getUserList)           // 获取列表
}
