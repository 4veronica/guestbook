package API

import (
	"github.com/gin-gonic/gin"
	"guestbook/database"
	"guestbook/model"
	"strconv"
)

func ReadPage(c *gin.Context) {
	clientID := c.Query("clientID")
	pageNumberFromClient := c.Query("pageNumber")

	pageNumber, err := strconv.Atoi(pageNumberFromClient)
	if err != nil {
		// ... handle error
		panic(err)
	}
	webClient := model.Client{ID: clientID}
	pageInfo := database.LoadFromDB(pageNumber)
	responseReadPage(c, pageInfo, webClient)
}

func responseReadPage(c *gin.Context, pageInfo model.Page, webclient model.Client) {
	c.JSON(200, gin.H{
		"clientID":   webclient.ID,
		"pageNumber": pageInfo.PageID,
	})
}
