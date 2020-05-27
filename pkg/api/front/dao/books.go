package dao

import (
	"lime/pkg/api/admin/model"
	"lime/pkg/common/db"
)

type BooksDao struct {}

func (c BooksDao) Get(id int) model.Books {
	var Books model.Books
	db := db.GetGormDB()
	db.Where("id = ?", id).First(&Books)
	return Books
}

func (c BooksDao) GetRandBooks(extraId int) []model.Books {
	var Books []model.Books
	db := db.GetGormDB()
	db.Model(&model.Books{}).Where("id !=?",extraId).Order("created_at desc").Limit(2).Find(&Books)
	return Books
}