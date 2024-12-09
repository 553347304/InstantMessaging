package mtype

import (
	"database/sql/driver"
	"encoding/json"
)

type SystemMessage struct {
	Type int8 `json:"type"` // 1 涉黄 2 社恐 3 涉政 4 不正当言论
}

func (c *SystemMessage) Scan(value interface{}) error { return json.Unmarshal(value.([]byte), c) }
func (c SystemMessage) Value() (driver.Value, error)  { return json.Marshal(c) }
