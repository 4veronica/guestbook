package database

import "fmt"

func DeleteToDB(clientID string, pageID int) error {
	var isPossibleDBConnect bool
	isPossibleDBConnect = true

	if !isPossibleDBConnect {
		return fmt.Errorf("DB 연결이 불가능 합니다.")
	} else {
		//PageID를 DB에 넘겨주어 삭제 한다.
		_, _ = clientID, pageID
		return nil
	}
}
