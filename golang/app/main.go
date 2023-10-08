package main

import (
	"context"
	"log"
	"server/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	entClient, err := middleware.NewEntClient()
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer entClient.Close()
	log.Println("DB Connected!")

	if err := entClient.Migrate(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	log.Println("DB Schema Created!")

}
