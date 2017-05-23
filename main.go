package main

import (
	"fmt"
	"log"
	"net/http"
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

	r.GET("/home", home)
	logger.Printf("Ready to listening the port %q..", port)

	r.Listen(":" + port)
}

func home(c *router.Control) {
	log.Print("Home")
	fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Request.URL.Path)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
