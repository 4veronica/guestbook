package API

import (
	"github.com/gin-gonic/gin"
	"guestbook/database"
	"net/http"
)

func UpdatePage(c *gin.Context) {
	pageInfo, err := getClientPage(c)
	if err == nil {
		err = database.WriteChangeToDB(pageInfo)
		if err == nil {
			c.JSON(http.StatusOK, "Page 수정 완료")
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}
