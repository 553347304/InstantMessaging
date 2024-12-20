package models

import (
	"database/sql/driver"
	"encoding/json"
)

type VerifyInfo struct {
	Issue  StringArr `json:"issue,omitempty"`
	Answer StringArr `json:"answer,omitempty"`
}

func (v VerifyInfo) Value() (driver.Value, error)  { return json.Marshal(v) }
func (v *VerifyInfo) Scan(value interface{}) error { return json.Unmarshal(value.([]byte), v) }

func (v VerifyInfo) Verify(v1 []string)bool  {
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
	// ID           uint                 `json:"id"`
	// String       string               `json:"string"`
	// AuthQuestion AuthQuestion         `json:"auth_question,omitempty"`
	// Answer       StringArr `json:"answer"`
}
