package main

import (
	fiber "github.com/gofiber/fiber/v2"
	//"github.com/matthewzhaocc/common"
	//log "github.com/sirupsen/logrus"
)

type config struct {
	record           string
	bearer           string
	bearercookiename string
}

func init() {

}

func main() {
	app := fiber.New()
	app.Use("/api/route", func(c *fiber.Ctx) error {
		return nil
	})
}
