package storage

import (
	todomodel "Project/module/item/model"
	"context"
)

func (s *sqlStorage) NewCreateItem(ctx context.Context, data *todomodel.ToDoItem) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
