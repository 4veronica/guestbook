package database

import (
	"fmt"
	"guestbook/model"
)

func WriteToDB(pageInfo *model.Page) error {
	pageID := pageInfo.PageID
	pageTitle := pageInfo.Title
	pageContents := pageInfo.Content
	pageTime := pageInfo.MadeTime

	var isPossibleDBConnect bool
	isPossibleDBConnect = true

	//DB 연결 가능여부 확인 후 불가능하면 error를 가능하면 nil을 리턴
	if !isPossibleDBConnect {
		return fmt.Errorf("DB 연결이 불가능 합니다.")
	} else {
		//사용자로 부터 받은 아래의 정보를 DB에 넘겨줘서 저장한다.
		_, _, _, _ = pageID, pageTitle, pageContents, pageTime
		return nil
	}
}
