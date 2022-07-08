package database

import (
	"fmt"
	"guestbook/model"
	"time"
)

func WriteChangeToDB(pageID int, pageTitle, pageContent string) error {
	var isPossibleDBConnect bool
	isPossibleDBConnect = true

	if !isPossibleDBConnect {
		return fmt.Errorf("DB 연결이 불가능 합니다.")
	} else {
		//json을 map으로 받는 방법 찾다가 그냥 포기하고 슬라이스로 - _ -
		targetIndex := -1
		for i := 0; i < len(model.QuestBookDB); i++ {
			if model.QuestBookDB[i].PageID == pageID {
				targetIndex = i
			}
		}

		if targetIndex == -1 {
			return fmt.Errorf("해당 페이지는 존재하지 않습니다.")
		} else {
			model.QuestBookDB[targetIndex].Title = pageTitle
			model.QuestBookDB[targetIndex].MadeTime = time.Now()
			model.QuestBookDB[targetIndex].Content = pageContent
		}
		return nil
	}
}
