package models

import (
	"database/sql/driver"
	"encoding/json"
)

type SystemMessage struct {
	Type int8 `json:"type"` // 1 涉黄 2 社恐 3 涉政 4 不正当言论
}

// Scan 取出来的时候的数据
func (c *SystemMessage) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), c)
}

// Value 入库的数据
func (c SystemMessage) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}
