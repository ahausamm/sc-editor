package main

import (
	"./application/backend"
	"fmt"
	"flag"
)

func main() {
	userId := flag.String("userId", "", "a string with the userId")
	flag.Parse()

	if(*userId != "") {
		api.Init(*userId)
	} else {
		fmt.Println("Error: please add a userId")
	}
}
