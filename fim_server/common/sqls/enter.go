package sqls

import "fmt"

func InNumber(field string, Id uint64) string {
	id1 := fmt.Sprintf("%%%d,%%", Id)
	id2 := fmt.Sprintf("%%%d,%%", Id)
	return fmt.Sprintf("(%s not like %s or %s not like %s)",
		field, field, id1, id2)
}
