package model

import "errors"

type Template struct {
	Id       int    `json:"template_id" db:"id"`
	OwnerId  string `json:"owner_id" db:"owner_id" binding:"required"`
	Name     string `json:"name" db:"name" binding:"required"`
	FileName string `json:"filename" db:"filename" binding:"required"`
}

type NewTemplate struct {
	Name     string `json:"name" db:"name" binding:"required"`
	FileName string `json:"filename" db:"filename" binding:"required"`
	Text     string `json:"text"`
}

type Content struct {
	Text string `json:"text"`
}

type EditTemplate struct {
	FileName string `json:"filename"`
	Text     string `json:"text"`
}

type UpdateListInput struct {
	Name     *string `json:"name"`
	FileName *string `json:"filename"`
}

func (i UpdateListInput) Validate() error {
	if i.Name == nil && i.FileName == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
