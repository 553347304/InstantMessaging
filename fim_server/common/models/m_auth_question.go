package models

import (
	"database/sql/driver"
	"encoding/json"
)

type AuthQuestion struct {
	Problem1 *string `json:"problem1"`
	Problem2 *string `json:"problem2"`
	Problem3 *string `json:"problem3"`
	Answer1  *string `json:"answer1"`
	Answer2  *string `json:"answer2"`
	Answer3  *string `json:"answer3"`
}

// Scan 取出来的时候的数据
func (c *AuthQuestion) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), c)
}

// Value 入库的数据
func (c *AuthQuestion) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
