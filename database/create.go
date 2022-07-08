package database

import (
	"fmt"
	"guestbook/model"
	"time"
)

func WriteToDB(title, content string) error {
	pageTitle := title
	pageContents := content

	var isPossibleDBConnect bool
	isPossibleDBConnect = true

	//DB 연결 가능여부 확인 후 불가능하면 error를 가능하면 nil을 리턴
	if !isPossibleDBConnect {
		return fmt.Errorf("DB 연결이 불가능 합니다.")
	} else {
		//사용자로 부터 받은 아래의 정보를 DB에 넘겨줘서 저장한다
		var pageID int
		if len(model.QuestBookDB) == 0 {
			pageID = 0
		} else {
			pageID = model.QuestBookDB[len(model.QuestBookDB)-1].PageID + 1
		}
		madeTime := time.Now()
		model.QuestBookDB = append(model.QuestBookDB, model.Page{PageID: pageID, Title: pageTitle, Content: pageContents, MadeTime: madeTime})
		return nil
	}
}
