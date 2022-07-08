package database

import (
	"guestbook/model"
	"time"
)

func LoadFromDB(pageNumber int) (pageInfo model.Page) {
	//DB를 연결해서 페이지 정보를 받아 리턴한다.
	pageInfo = model.Page{
		PageID:   pageNumber,
		Title:    "방명록 연습중",
		Content:  "방명록 서버를 만들면서 공부 하고 있습니다",
		MadeTime: time.Now(),
	}
	return pageInfo
}
