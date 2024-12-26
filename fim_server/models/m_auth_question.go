package models

import (
	"database/sql/driver"
	"encoding/json"
	"fim_server/models/mgorm"
)

type ValidInfo struct {
	Issue  mgorm.String `json:"issue,omitempty"`
	Answer mgorm.String `json:"answer,omitempty"`
}

func (v ValidInfo) Value() (driver.Value, error)  { return json.Marshal(v) }
func (v *ValidInfo) Scan(value interface{}) error { return json.Unmarshal(value.([]byte), v) }

func (v ValidInfo) Valid(v1 []string) bool {
	if len(v.Answer) == len(v1) {
		for i, _ := range v.Answer {
			if v.Answer[i] != v1[i] {
				return false
			}
		}
		return true
	}
	return false
}

type Test struct {
	ID     uint         `json:"id"`
	String string       `json:"string"`
	Arr    mgorm.String `json:"arr"`
}
