package repository

import (
	"github.com/jmoiron/sqlx"
	"template-server/model"
)

type TemplateManager interface {
	CreateTemplate(userId int, template model.NewTemplate) (int, error)
	DeleteTemplate(templateId int) error
	GetTemplate(templateId int) (model.Template, error)
	GetAll(userId int) ([]model.Template, error)
	UpdateTemplate(userId int, id int, input model.EditTemplate) interface{}
}

type Repository struct {
	TemplateManager
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TemplateManager: NewTemplateBD(db),
	}
}
