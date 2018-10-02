package main

import (
	"fmt"

	"github.com/kodywilson/callistaenterprise/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
}
