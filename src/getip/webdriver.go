/*
	Copyright Matthew 2021
	getip the web API that gets the requester IP
*/
package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/matthewzhaocc/common"
)

// WhatIsMyIP gets the actual IP function
func WhatIsMyIP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Host)

	log.Println("Received Request from ", r.Host)
}
// main handler
func main() {
	log.Println("starting webserver for the GetMyIP Service")
	http.HandleFunc("/api/ip", WhatIsMyIP)
	log.Println("Starting service at port :6443")
	http.ListenAndServe(":6443", nil)
}
