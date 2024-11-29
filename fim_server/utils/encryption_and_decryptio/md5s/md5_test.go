package md5s

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	md5 := Hash([]byte("123456"))
	fmt.Println(md5)
}
