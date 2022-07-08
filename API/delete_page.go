package API

import (
	"github.com/gin-gonic/gin"
	"guestbook/database"
	"net/http"
	"strconv"
)

func DeletePage(c *gin.Context) {
	clientID := c.Query("clientID") // shortcut for c.Request.URL.Query().Get("lastname")
	pageNumberFromClient := c.Query("pageNumber")
	pageNumber, err := strconv.Atoi(pageNumberFromClient)
	if err == nil {
		err = database.DeleteToDB(clientID, pageNumber)
		if err == nil {
			c.JSON(http.StatusOK, "Page 삭제")
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}
