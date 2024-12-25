package main

import (
	"fim_server/utils/stores/method"
	"fmt"
)

func main() {
	// 定义一个整数切片
	// numbers := []int{9, 8, 1, 3, 5,6 ,7}
	numbers := []string{"9", "8", "1", "3", "5","6" ,"7"}


	fmt.Println(method.List(numbers).Sort(true)) // 输出: [1 2 3 4 7]



}



