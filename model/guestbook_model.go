package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type Page struct {
	PageID   int       `json:"pageID"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	MadeTime time.Time `json:"madeTime"`
}

func (b *Page) UnmarshalJSON(data []byte) error {
	type Alias Page
	clientJson := &struct{ *Alias }{Alias: (*Alias)(b)}
	if err := json.Unmarshal(data, &clientJson); err != nil {
		return err
	}
	fmt.Println(clientJson)
	return nil
}
