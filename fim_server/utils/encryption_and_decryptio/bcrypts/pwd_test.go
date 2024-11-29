package bcrypts

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	hash := Hash("by")
	fmt.Println(hash)
}

func TestCheckPwd(t *testing.T) {
	//
	// $2a$04$V4VS9vHKN0.4Hzc5k2fKSeRk1K8Ui1ubFZUVrYaFoym3..jnDpkTK
	ok := Check("$2a$04$kR0y.JY48JwxpWgHWPJkuuQkBFa0hpKMLVEKg3diTHHHK7kHaVkZa", "by")
	fmt.Println(ok)
}
