package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	"template-server/model"
)

const (
	templateTable = "templates"
)

type TemplateDB struct {
	db *sqlx.DB
}

func NewTemplateBD(db *sqlx.DB) *TemplateDB {
	return &TemplateDB{db: db}
}

func (r *TemplateDB) CreateTemplate(userId int, template model.NewTemplate) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (owner_id, name, filename) values ($1, $2, $3) RETURNING id", templateTable)

	row := r.db.QueryRow(query, userId, template.Name, template.FileName)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TemplateDB) DeleteTemplate(templateId int) error {
	query := fmt.Sprintf("DELETE FROM %s "+
		"WHERE id = $1", templateTable)

	_, err := r.db.Exec(query, templateId)
	return err
}

func (r *TemplateDB) GetTemplate(templateId int) (model.Template, error) {
	var item model.Template
	query := fmt.Sprintf(`SELECT ti.id, ti.owner_id, ti.name, ti.filename FROM %s ti 
								WHERE ti.id = $1 `,
		templateTable)
	if err := r.db.Get(&item, query, templateId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TemplateDB) GetAll(userId int) ([]model.Template, error) {
	var items []model.Template
	query := fmt.Sprintf(`SELECT ti.id, ti.owner_id, ti.name, ti.filename FROM %s ti 
								WHERE ti.owner_id = $1 `,
		templateTable)
	if err := r.db.Select(&items, query, userId); err != nil {
		return items, err
	}

	return items, nil
}

func (r *TemplateDB) UpdateTemplate(userId int, id int, input model.EditTemplate) interface{} {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Text != "" {
		setValues = append(setValues, fmt.Sprintf("text=$%d", argId))
		args = append(args, input.Text)
		argId++
	}
	args = append(args, id)

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d",
		templateTable, setQuery, argId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("user: %d", userId)

	_, err := r.db.Exec(query, args...)
	return err
}
