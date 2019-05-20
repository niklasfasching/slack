package main

import (
	"log"
	"os"

	"github.com/niklasfasching/slack"
)

func main() {
	c := &slack.Connection{
		Token: os.Getenv("slack_token"),
		Debug: true,
	}
	if err := c.Start(nil); err != nil {
		log.Println(err)
	}
}
