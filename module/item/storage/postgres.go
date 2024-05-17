package storage

import (
	todomodel "Project/module/item/model"
	"context"
	"gorm.io/gorm"
)

type sqlStorage struct {
	db *gorm.DB
}

func (s *sqlStorage) ListItem(ctx context.Context, condition map[string]interface{}, paging *todomodel.DataPaging) ([]todomodel.ToDoItem, error) {
	//TODO implement me
	panic("implement me")
}

func NewSQLStorage(db *gorm.DB) *sqlStorage {
	return &sqlStorage{db: db}
}
