package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Model struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ContentArray []Content

func (s ContentArray) Value() (driver.Value, error) {
	if s == nil {
		return "[]", nil
	}
	return json.Marshal(s)
}
func (s *ContentArray) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	if len(bytes) > 0 {
		return json.Unmarshal(bytes, s)
	}
	*s = make(ContentArray, 0)
	return nil
}

type Test struct {
	ID  uint         `json:"id"`
	Arr ContentArray `json:"arr"`
}

type Content struct {
	Type    int8   `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Size    int    `json:"size,omitempty"`
}

