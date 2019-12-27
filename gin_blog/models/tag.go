package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name string `json:"name"`
	CreatedBy string `json:"create_by"`
	ModifiedBy string `json:"modifide_by"`
	State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag)  {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).Find(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}
func ExistTagById(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).Find(&tag)
	if tag.ID > 0 {
		return true
	}

	return false
}

func AddTag(name string, state int, createdBy string) int {
	tag := Tag{
		Name:       name,
		CreatedBy:  createdBy,
		State:      state,
	}
	db.Create(&tag)
	return tag.ID
}

func Edit(id int, data interface{}) bool  {
	db.Model(&Tag{}).Where("id = ?", id).Update(data)
	return true
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func (tag *Tag)BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (tag *Tag)BeforeUpdate(scope *gorm.Scope) error {
	return scope.SetColumn("ModifiedOn", time.Now().Unix())
}