package API

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"guestbook/database"
	"guestbook/model"
	"net/http"
)

func CreatePage(c *gin.Context) {
	pageInfo, err := getClientPage(c)
	if err == nil {
		err = database.WriteToDB(pageInfo)
		if err == nil {
			c.JSON(http.StatusOK, "Page 저장이 완료 되었음")
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	} else {
		c.JSON(http.StatusBadRequest, err.Error())
	}
}

func getClientPage(c *gin.Context) (*model.Page, error) {
	var clientPage model.Page
	err := c.Bind(&clientPage)
	if err != nil {
		return nil, fmt.Errorf("사용자 입력정보 에러")
	} else {
		return &clientPage, nil
	}
}
