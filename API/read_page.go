package API

import (
	"github.com/gin-gonic/gin"
	"guestbook/database"
	"guestbook/model"
	"net/http"
)

func ReadPage(c *gin.Context) {
	allPages := database.LoadAllPagesFromDB()
	responseReadPage(c, allPages)
}

func responseReadPage(c *gin.Context, allPages []model.Page) {
	if len(allPages) == 0 {
		c.JSON(http.StatusOK, "방명록이 존재 하지 않습니다.")
	}

	for i := 0; i < len(allPages); i++ {
		c.JSON(200, gin.H{
			"Page ID":    allPages[i].PageID,
			"Page Title": allPages[i].Title,
			//		"Page Content":   allPages[i].Content,
			//		"Page Made Time": allPages[i].MadeTime,
		})
	}
}
