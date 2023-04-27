package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sbuttigieg/test-xm/cmd/config"
	"github.com/sbuttigieg/test-xm/cmd/config/connections"
	"github.com/sbuttigieg/test-xm/cmd/config/store"
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

	// store setup
	err = store.DBInit(dbConnection)
	if err != nil {
		// log.Error("Database Initiation Error: ", err)
		fmt.Println("Database Initiation Error: ", err)
	}

	fmt.Println("test-xm")
}
