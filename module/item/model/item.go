package todomodel

import "time"

type ToDoItem struct {
	Id        int        `json:"id" gorm:"column:id"`
	Name      string     `json:"name" gorm:"column:name"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (ToDoItem) TableName() string { return "todo_items" }

func (item ToDoItem) Validate() error {
	return nil
}

type DataPaging struct {
	Page  int   `json:"page" form:"page"`
	Limit int   `json:"limit" form:"limit"`
	Total int64 `json:"total" form:"-"`
}

func (p *DataPaging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}
}
