package API

import (
	"github.com/gin-gonic/gin"
	"guestbook/database"
	"net/http"
	"strconv"
)

func DeletePage(c *gin.Context) {
	pageIDFromClient := c.Query("pageID")
	pageID, err := strconv.Atoi(pageIDFromClient)
	if err == nil {
		err = database.DeleteToDB(pageID)
		if err == nil {
			c.JSON(http.StatusOK, "Page 삭제")
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}
