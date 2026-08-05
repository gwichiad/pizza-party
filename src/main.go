package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main()  {

	ctx := context.Background()
	mongo_uri := "mongodb://admin:password@localhost:27017"

	mongoClient, err := startDB(ctx, mongo_uri)
	if err != nil {
		panic(err)
	}
	defer func() {
		disconnectCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := mongoClient.Disconnect(disconnectCtx); err != nil {
			panic(err)
		}
		fmt.Println("Disconnected from MongoDB")
	}()

	data := newStore(mongoClient)

	cfg := config{
		addr: ":8080",
		store: data,
	}
	api := api{
		config: cfg,
	}

	go api.run(api.mount())

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down")
}
