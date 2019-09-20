package main

import (
	"fmt"
	"log"
	"net/http"
	"whacos/pkg/settings"
	"whacos/routers"
)

// @title Swagger For Project Whacos
// @version 1.0
// @description This is a api server for project whacos .
// @termsOfService http://swagger.io/terms/
// @contact.name API 接口
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 148.70.20.37:8090
// @BasePath /
func main() {

	router := routers.InitRouter()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", settings.HttpPort),
		Handler:        router,
		ReadTimeout:    settings.ReadTimeout,
		WriteTimeout:   settings.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	/**
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
	*/
}
