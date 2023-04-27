package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sbuttigieg/test-xm/cmd/config"
	"github.com/sbuttigieg/test-xm/cmd/config/connections"
)

func main() {
	ctx := context.Background()

	// config
	c, err := config.NewConfig()
	if err != nil {
		log.Fatal(ctx, err.Error())
	}

	// connections
	dbConnection, err := connections.NewPostgres(c)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(dbConnection, "test-xm")
}
