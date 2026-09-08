package config

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServer(){
	cfg := LoadAppConfig("")

	port := cfg.AppPort

	srv := &http.Server{
		Addr: port,
		//Place Router to be used as handler in app
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout: 6 * time.Second,
	}

	go func(){
		fmt.Printf("Server started %v application listening at %v\n", cfg.AppName, cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Fatal Error, Failed to listen at %v, : %v\n", cfg.AppPort, err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	fmt.Println("Shutting Server Gracefully...")

	//Here you can stop any existing background cron jo
		//eg alertCronJob.Stop()
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Printf("%v stopped/shutdown properly.\n", cfg.AppName)

}