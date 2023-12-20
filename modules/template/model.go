package template

type TemplateModel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Num       int    `json:"num"`
	IsDeleted bool   `json:"isDeleted"`
}

type PageDTO struct {
	Page     int `query:"page" validate:"required,gte=1" example:"1"`
	PageSize int `query:"pageSize" validate:"required,gte=1" example:"10"`
}

type IdDTO struct {
	ID int `param:"id" validate:"required" example:"1"`
}
