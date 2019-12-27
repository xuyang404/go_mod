package models

type Article struct {
	Model
	TagId int `json:"tag_id" gorm:"index"`
}