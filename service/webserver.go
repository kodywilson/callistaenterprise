package service

import (
	"log"
	"net/http"
)

// Start listening for http requests on passed port
func StartWebServer(port string) {

	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here

	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}

}
