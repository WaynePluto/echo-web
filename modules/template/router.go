package template

import "github.com/labstack/echo/v4"

func SetRouter(e *echo.Echo) {
	e.POST("/template", addTemplate)              // 创建
	e.GET("/template/:id", getTemplateById)       // 获取单个
	e.PUT("/template/:id", updateTemplateById)    // 修改
	e.PATCH("/template/:id", updateTemplateById)  // 部分修改
	e.DELETE("/template/:id", deleteTemplateById) // 删除
	e.GET("/template", getTemplateList)           // 获取列表
}
