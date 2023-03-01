package service

import (
	"template-server/api/repository"
	"template-server/model"
)

type Service struct {
	TemplateItem
}

type TemplateItem interface {
	Create(userId int, template model.NewTemplate) (int, error)
	GetAll(userId int) ([]model.Template, error)
	GetById(userId int, templateId int) (model.Template, error)
	Delete(templateId int) error
	Update(input model.EditTemplate) interface{}
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TemplateItem: NewTemplateService(repos.TemplateManager),
	}
}
