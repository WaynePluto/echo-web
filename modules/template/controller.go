package template

import (
	"echo-web/custom"

	"github.com/labstack/echo/v4"
)

// addTemplate godoc
//
//	@Summary			addTemplate
//	@Description
//	@Tags					template
//	@Accept				json
//	@Produce			json
//	@Param				template-info	body	TemplateModel	true	"template info object"
//	@Success			200 {object} TemplateModel
//	@Failure			400 {object} interface{}
//	@Router				/templates [post]
func addTemplate(c echo.Context) error {
	cc := c.(*custom.Context)
	var template TemplateModel
	err := cc.BindDTO(&template)
	if err != nil {
		return cc.ClientErr(err)
	}
	cc.Logger().Infof("template dto: %v", template)
	// service

	return cc.Send(template)
}

// getTemplateById godoc
//
//	@Summary		getTemplateById
//	@Description
//	@Tags				template
//	@Accept			json
//	@Produce		json
//	@Param			id	path	string	true	"template ID"
//	@Success		200 {object} TemplateModel
//	@Failure		400 {object} interface{}
//	@Router			/templates/{id} [get]
func getTemplateById(c echo.Context) error {
	cc := c.(*custom.Context)
	var idDTO IdDTO
	cc.BindDTO(&idDTO)
	cc.Logger().Infof("template id: %v", idDTO)
	return cc.Send(TemplateModel{Num: idDTO.ID})
}

// updateTemplateById godoc
//
//	@Summary		updateTemplateById
//	@Description
//	@Tags				template
//	@Accept			json
//	@Produce		json
//	@Param			id	path	number	true	"template ID"
//	@Success		200 {object} TemplateModel
//	@Failure		400 {object} interface{}
//	@Router			/templates/{id} [put]
func updateTemplateById(c echo.Context) error {
	cc := c.(*custom.Context)
	return cc.Send(TemplateModel{})
}

// updateTemplateById godoc
//
//	@Summary		updateTemplateById
//	@Description
//	@Tags				template
//	@Accept			json
//	@Produce		json
//	@Param			id	path	number	true	"template ID"
//	@Success		200 {object} TemplateModel
//	@Failure		400 {object} interface{}
//	@Router			/templates/{id} [delete]
func deleteTemplateById(c echo.Context) error {
	cc := c.(*custom.Context)
	return cc.Send(TemplateModel{})
}

// getTemplateList godoc
//
//	@Summary			getTemplateList
//	@Description
//	@Tags					template
//	@Accept				json
//	@Produce			json
//	@Param				page	query	PageDTO	true	"page"
//	@Success			200 {object} []TemplateModel
//	@Failure			400 {object} interface{}
//	@Router				/templates [get]
func getTemplateList(c echo.Context) error {
	cc := c.(*custom.Context)

	var pageDTO PageDTO
	err := cc.BindDTO(&pageDTO)
	if err != nil {
		return cc.ClientErr(err)
	}

	return cc.Send([]TemplateModel{{Num: pageDTO.Page}})
}
