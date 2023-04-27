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
	dbConnection, err := connections.NewMySQL(c)
	if err != nil {
		log.Fatal(err.Error())
	}

	dbConnection2, err := connections.NewPostgres(c)
	if err != nil {
		log.Fatal(err.Error())
	}

	// store setup
	err = store.DBInit(dbConnection)
	if err != nil {
		// log.Error("Database Initiation Error: ", err)
		fmt.Println("Database Initiation Error: ", err)
	}

	err = store.DBInit2(dbConnection2)
	if err != nil {
		// log.Error("Database Initiation Error: ", err)
		fmt.Println("Database Initiation Error: ", err)
	}

	fmt.Println("test-xm")

	// const errorChan int = 10
	// errChan := make(chan error, errorChan)
	// go func() {
	// 	fmt.Println("*****", "go-1")
	// 	// log.Info(ctx, fmt.Sprintf("Health service listening on %s", healthAddr))
	// 	// errChan <- healthServer.ListenAndServe()
	// }()
	// go func() {
	// 	fmt.Println("*****", "go-2")
	// 	// log.Info(ctx, fmt.Sprintf("HTTP service listening on %s", httpAddr))
	// 	// errChan <- httpServer.ListenAndServe()
	// }()

	// signalChan := make(chan os.Signal, 1)
	// signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// for {
	// 	select {
	// 	case err := <-errChan:
	// 		if err != nil {
	// 			log.Fatal(ctx, err.Error())
	// 		}
	// 	case s := <-signalChan:
	// 		fmt.Println("***s**", s)
	// 		// log.Info(ctx, "Captured %v. Exiting...", s)
	// 		// healthserv.SetReadinessStatus(http.StatusServiceUnavailable)
	// 		// httpServer.BlockingClose()
	// 		os.Exit(0)
	// 	}
	// }
}
