package database

import (
	"fmt"
	"guestbook/model"
)

func WriteChangeToDB(pageInfo *model.Page) error {
	var isPossibleDBConnect bool
	isPossibleDBConnect = true

	if !isPossibleDBConnect {
		return fmt.Errorf("DB 연결이 불가능 합니다.")
	} else {
		//사용자로 부터 받은 아래의 정보를 DB에 넘겨줘서 저장한다.
		_ = pageInfo
		return nil
	}
}
