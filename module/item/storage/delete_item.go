package storage

import (
	todomodel "Project/module/item/model"
	"context"
)

func (s *sqlStorage) DeleteItem(
	ctx context.Context,
	condition map[string]interface{},
) error {

	if err := s.db.
		Table(todomodel.ToDoItem{}.TableName()).
		Where(condition).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
