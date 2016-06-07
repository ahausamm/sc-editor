package main

import (
	"./application/backend"
	"fmt"
	"flag"
)

func main() {
	UserId := flag.String("UserId", "", "a string with the UserId")
	InstanceId := flag.String("InstanceId", "", "a string with the InstanceId")
	ServiceUrl := flag.String("ServiceUrl", "", "a string with the ServiceUrl")
	ApiVersion := flag.String("ApiVersion", "", "a string with the ApiVersion")
	flag.Parse()
	if(*UserId != "" && *InstanceId != "" && *ServiceUrl != "" && *ApiVersion != "") {
		api.Init(*UserId,*InstanceId,*ServiceUrl,*ApiVersion)
	} else {
		fmt.Println("Error: please add a UserId, a InstanceId, a ServiceUrl and an ApiVersion")
	}
}
