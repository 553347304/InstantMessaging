package mgorm

import (
	"database/sql/driver"
	"encoding/json"
)

type Uint64 []uint64
type Int []int
type String []string

func (s String) Value() (driver.Value, error) {
	if s == nil {
		return "[]", nil
	}
	return json.Marshal(s)
}
func (s *String) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	if len(bytes) > 0 {
		return json.Unmarshal(bytes, s)
	}
	*s = make(String, 0)
	return nil
}
func (s Int) Value() (driver.Value, error) {
	if s == nil {
		return "[]", nil
	}
	return json.Marshal(s)
}
func (s *Int) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	if len(bytes) > 0 {
		return json.Unmarshal(bytes, s)
	}
	*s = make(Int, 0)
	return nil
}
func (s Uint64) Value() (driver.Value, error) {
	if s == nil {
		return "[]", nil
	}
	return json.Marshal(s)
}
func (s *Uint64) Scan(value interface{}) error {
	bytes, _ := value.([]byte)
	if len(bytes) > 0 {
		return json.Unmarshal(bytes, s)
	}
	*s = make(Uint64, 0)
	return nil
}

