package main

import (
	"fim_server/config"
	"fim_server/models"
	"fmt"
)

type Content struct {
	Type    int8   `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type Size struct {
	Type    int8   `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Size    string `json:"size"`
}

func main() {

	config.Init()

	var scan *[]models.Test

	find(&scan)

	// for _, test := range scan {
	// 	fmt.Println("t",test)
	// }
	fmt.Println(scan)

}

func find(scan interface{}) {

	config.DB.Find(scan)
}
