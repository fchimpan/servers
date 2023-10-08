package app

import (
	"fmt"
	"log"
	"os"

	"ariga.io/sqlcomment/examples/ent"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")
	if dbUser == "" || dbPass == "" || dbPort == "" || dbName == "" {
		log.Fatal("DB Config Error")
	}

	client, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(localhost:%s)/%s?parseTime=True", dbUser, dbPass, dbPort, dbName))
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	log.Println("DB Connected")
}
