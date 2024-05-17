package todobiz

import (
	todomodel "Project/module/item/model"
	"context"
)

type CreateTodoItemStorage interface {
	CreateItem(ctx context.Context, data *todomodel.ToDoItem) error
}

type createBiz struct {
	store CreateTodoItemStorage
}

func NewCreateToDoItemBiz(store CreateTodoItemStorage) *createBiz {
	return &createBiz{store: store}
}

func (biz *createBiz) CreateNewItem(ctx context.Context, data *todomodel.ToDoItem) error {
	if err := data.Validate(); err != nil {
		return err
	}

	// do not allow "finished" status when creating a new task
	data.Name = "Doing" // set to default

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
