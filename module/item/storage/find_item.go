package storage

import (
	todomodel "Project/module/item/model"
	"context"
)

func (s *sqlStorage) FindItem(
	ctx context.Context,
	condition map[string]interface{},
) (*todomodel.ToDoItem, error) {
	var itemData todomodel.ToDoItem

	if err := s.db.Where(condition).First(&itemData).Error; err != nil {
		return nil, err // other errors
	}

	return &itemData, nil
}
