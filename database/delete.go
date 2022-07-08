package database

import (
	"fmt"
	"guestbook/model"
)

func DeleteToDB(pageID int) error {
	var isPossibleDBConnect bool
	isPossibleDBConnect = true

	if !isPossibleDBConnect {
		return fmt.Errorf("DB 연결이 불가능 합니다.")
	} else {
		targetIndex := -1
		for i := 0; i < len(model.QuestBookDB); i++ {
			if model.QuestBookDB[i].PageID == pageID {
				targetIndex = i
			}
		}
		if targetIndex == -1 {
			return fmt.Errorf("해당 페이지는 존재하지 않습니다.")
		} else {
			model.QuestBookDB = append(model.QuestBookDB[:targetIndex], model.QuestBookDB[targetIndex+1:]...)
		}
		return nil
	}
}
