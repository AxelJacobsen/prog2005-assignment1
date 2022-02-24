package main

import (
	RESTuni "cloudAss1"
	"net/http"
	"time"
)

// -----------------
func main() {
	dt := time.Now().Unix()
	/*
		port := os.Getenv("PORT")
		if port == "" {
			log.Println("$PORT has not been set. Default:" + RESTuni.LOCAL_PORT)
			port = RESTuni.LOCAL_PORT
		} */

	// Handles getting information from universites
	http.HandleFunc("/"+RESTuni.UNI_BASE_PATH+"/", RESTuni.HandlerUniversity(dt)) // ensure to type complete URL when requesting

	/* 	log.Println("Listening on port " + port)
	   	log.Fatal(http.ListenAndServe(":"+port, nil)) */
}
