package custom

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
}

func InitCtx(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &Context{c}
		return next(cc)
	}
}

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func (c *Context) Send(data any) error {
	r := Response{http.StatusOK, data, "success"}
	return c.JSON(http.StatusOK, r)
}

func (c *Context) ClientErr(e error) error {
	r := Response{http.StatusBadRequest, nil, e.Error()}
	return c.JSON(http.StatusOK, r)
}

func (c *Context) ServerErr(e error) error {
	r := Response{http.StatusInternalServerError, nil, e.Error()}
	return c.JSON(http.StatusOK, r)
}

// 数据绑定与校验
func (c *Context) BindDTO(data any) error {
	err := c.Bind(data)
	if err != nil {
		return err
	}
	err = c.Validate(data)
	if err != nil {
		return err
	}
	return nil
}
