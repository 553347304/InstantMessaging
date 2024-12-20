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




type StringArr []string

func (s StringArr) Value() (driver.Value, error) {
	if s == nil {
		return "[]", nil
	}
	return json.Marshal(s)
}
func (s *StringArr) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	if len(bytes) > 0 {
		return json.Unmarshal(bytes, s)
	}
	*s = make([]string, 0)
	return nil
}
