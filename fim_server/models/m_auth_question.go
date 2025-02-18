package models

import (
	"database/sql/driver"
	"encoding/json"
)

type ValidInfo struct {
	Issue  string `json:"issue,omitempty"`
	Answer string `json:"answer,omitempty"`
}

func (v ValidInfo) Value() (driver.Value, error)  { return json.Marshal(v) }
func (v *ValidInfo) Scan(value interface{}) error { return json.Unmarshal(value.([]byte), v) }

func (v ValidInfo) Valid(answer string) bool {
	return v.Answer == answer
}
