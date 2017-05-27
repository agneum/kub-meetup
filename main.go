package main

import (
	"log"
	"os"

	"github.com/takama/router"
)

func main() {
	logger := log.New(os.Stdout, "[step-by-step]", log.LstdFlags)
	logger.Print("App is running..")

	port := os.Getenv("SERVICE_PORT")

	if len(port) == 0 {
		logger.Fatal("Specify port")
	}

	r := router.New()

	r.GET("/", home)
	logger.Printf("Ready to listening the port %q..", port)

	r.Listen(":" + port)
}
