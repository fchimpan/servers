package middleware

import (
	"context"
	"fmt"
	"os"
	"server/ent"
)

type EntClient struct {
	*ent.Client
}

func NewEntClient() (*EntClient, error) {
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")
	if dbUser == "" || dbPass == "" || dbPort == "" || dbName == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	client, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		return nil, err
	}

	if os.Getenv("ENV") == "dev" {
		client = client.Debug()
	}

	return &EntClient{client}, nil
}

func (c *EntClient) Migrate(ctx context.Context) error {
	if err := c.Schema.Create(ctx); err != nil {
		return err
	}
	return nil
}
