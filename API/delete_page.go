package API

import (
	"github.com/gin-gonic/gin"
	"guestbook/database"
	"net/http"
	"strconv"
)

func DeletePage(c *gin.Context) {
	pageNumberFromClient := c.Query("pageNumber")
	pageNumber, err := strconv.Atoi(pageNumberFromClient)
	if err == nil {
		err = database.DeleteToDB(pageNumber)
		if err == nil {
			c.JSON(http.StatusOK, "Page 삭제")
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}
