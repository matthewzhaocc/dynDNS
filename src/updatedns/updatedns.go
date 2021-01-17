package main

import (
	"os"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/matthewzhaocc/common"
	log "github.com/sirupsen/logrus"
)

func UpdateDNS(url string) {
	_, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if common.CheckError(err) {
		log.Println("error !!!!!" + err.Error())
	}

}
