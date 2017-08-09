package main

import (
	"fmt"
	"os"
	"time"

	"zxq.co/ripple/vinse/vinstore/models"
	"zxq.co/ripple/vinse/vinstore/webserver"
)

func main() {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		dsn = "root@/vinse"
	}
	db, err := models.Create(dsn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Listening on", ":51512")
	webserver.DB = db
	go webserver.Start(":51512")

	c := &Client{
		DB: db,
	}
	typeMap := c.GenerateTypeMap()

	for {
		err := ListenToWS(typeMap)
		if err != nil {
			fmt.Println("Error connecting to WebSocket:", err)
			fmt.Println("Retrying in 10 seconds...")
		}
		time.Sleep(time.Second * 10)
	}
}
