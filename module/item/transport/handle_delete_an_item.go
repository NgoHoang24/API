package todotrpt

import (
	todobiz "Project/module/item/biz"
	todostorage "Project/module/item/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func HandleDeleteAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// setup dependencies
		storage := todostorage.NewSQLStorage(db)
		biz := todobiz.NewDeleteToDoItemBiz(storage)

		if err := biz.DeleteItem(c.Request.Context(), map[string]interface{}{"id": id}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
