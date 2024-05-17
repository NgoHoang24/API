package todotrpt

import (
	todobiz "Project/module/item/biz"
	todostorage "Project/module/item/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleFindAnItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		storage := todostorage.NewSQLStorage(db)
		biz := todobiz.NewFindToDoItemBiz(storage)

		data, err := biz.FindAnItem(c.Request.Context(), map[string]interface{}{"id": id})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
	}
}
