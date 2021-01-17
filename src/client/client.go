// main the controller of the client
package main

import (
	"os"

	"github.com/matthewzhaocc/common"
)

type config struct {
	Server          string
	Bearer          string
	BearerTokenName string
}

var conf config

func init() {
	conf.Server = os.Getenv("SERVER_HOSTNAME")
	conf.Bearer = os.Getenv("BEARER")
	conf.BearerTokenName = os.Getenv("BEARER_TOKEN_NAME")
}

func main() {
	
}
