package bcrypts

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	hash := HashPwd("123456")
	fmt.Println(hash)
}

func TestCheckPwd(t *testing.T) {
	//
	// $2a$04$V4VS9vHKN0.4Hzc5k2fKSeRk1K8Ui1ubFZUVrYaFoym3..jnDpkTK
	ok := CheckPwd("$2a$04$RIqBGNbvUqc20b6qxL8KvuqEVXBqcs28E0ZJm.YBGV5S4JB7Ho.j2", "123456")
	fmt.Println(ok)
}
