package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"whacos/pkg/settings"
	"whacos/routers"
)

func main() {

	router := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HttpPort),
		Handler:        router,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
