package main

import (
	"./application/backend"
	"fmt"
)

func main() {

	api.Init()

	fmt.Println("run")
	api.Run()
}
