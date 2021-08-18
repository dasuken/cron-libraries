package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	c := cron.New()
	c.AddFunc("@every 2s", func() {
		fmt.Println("every 2s call")
	})

	c.AddFunc("@every 3s", func() {
		fmt.Println("every 3s call")
	})

	c.Start()

	time.Sleep(10 * time.Second)
}