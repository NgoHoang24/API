package todotrpt

import (
	todobiz "Project/module/item/biz"
	todomodel "Project/module/item/model"
	todostorage "Project/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func HandleCreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var dataItem todomodel.ToDoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// setup dependencies
		storage := todostorage.NewSQLStorage(db)
		biz := todobiz.NewCreateToDoItemBiz(storage)

		if err := biz.CreateNewItem(c.Request.Context(), &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": dataItem.Id})
	}
}
