package main

import (
	"fmt"
	"os"
	"time"

	"github.com/osukurikku/multis/vinstore/models"
	"github.com/osukurikku/multis/vinstore/webserver"
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

	fmt.Println("Listening on", "127.0.0.1:51512")
	webserver.DB = db
	go webserver.Start("127.0.0.1:51512")

	c := &Client{
		DB:       db,
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
