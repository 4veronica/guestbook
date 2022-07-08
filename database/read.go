package database

import (
	"guestbook/model"
)

func LoadAllPagesFromDB() []model.Page {
	//DB를 연결해서 페이지 정보를 받아 리턴한다.
	return model.QuestBookDB
}
