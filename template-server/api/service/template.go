package service

import (
	"template-server/api"
	"template-server/api/repository"
	"template-server/model"
)

type TemplateService struct {
	repo repository.TemplateManager
}

func NewTemplateService(repo repository.TemplateManager) *TemplateService {
	return &TemplateService{repo}
}

func (s *TemplateService) Create(userId int, template model.NewTemplate) (int, error) {
	err := api.CreateTemplateFile(template.FileName, template.Text)
	if err != nil {
		return 0, err
	}

	return s.repo.CreateTemplate(userId, template)
}

func (s *TemplateService) GetAll(userId int) ([]model.Template, error) {
	return s.repo.GetAll(userId)
}

func (s *TemplateService) GetById(userId int, templateId int) (model.Template, error) {
	return s.repo.GetTemplate(templateId)
}

func (s *TemplateService) Update(input model.EditTemplate) interface{} {
	/*loc, err := os.Getwd()
	if err != nil {
		return err
	}

	loc += viper.GetString("upload.location") + viper.GetString("upload.dir_name") + `\`

	err = ioutil.WriteFile(loc+input.FileName, []byte(input.Text), 0666)
	if err != nil {
		print(err.Error())
		return err
	}*/
	return nil
}

func (s *TemplateService) Delete(templateId int) error {
	return s.repo.DeleteTemplate(templateId)
}
