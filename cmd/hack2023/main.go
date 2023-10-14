package main

import (
	"hack2023/internal/app/config"
	"hack2023/internal/app/server"
	"hack2023/internal/app/store"
	"log"
	"net/http"
)

func main() {
	config := config.Get()

	defer func() {
		if msg := recover(); msg != nil {
			log.Println("Panic: ", msg)
		}
	}()

	//подключение к бд
	store, err := store.New(config)
	if err != nil {
		log.Print(err)
	}

	// ctx := context.Background()
	// opt := option.WithCredentialsFile(config.FireBaseFile)
	// app, err := firebase.NewApp(ctx, nil, opt)
	// if err != nil {
	// 	log.Print(err)
	// }
	// push, err := app.Messaging(ctx)
	// if err != nil {
	// 	log.Print(err)
	// }

	srv := server.NewServer(store, config)

	if err := srv.Start(config); err != nil && err != http.ErrServerClosed {
		log.Print(err)
	}

}
