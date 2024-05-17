package todotrpt

import (
	todobiz "Project/module/item/biz"
	todomodel "Project/module/item/model"
	todostorage "Project/module/item/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleUpdateAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var dataItem todomodel.ToDoItem

		if err := c.ShouldBind(&dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := todostorage.NewSQLStorage(db)
		biz := todobiz.NewUpdateToDoItemBiz(storage)

		if err := biz.UpdateItem(c.Request.Context(), map[string]interface{}{"id": id}, &dataItem); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
