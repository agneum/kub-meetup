package main

import (
	"fmt"
	"net/http"

	"github.com/takama/router"
)

func main() {
	r := router.New()

	a := 5
	// r.GET("/", home)
	// r.POST("/pampam", pampam)
	r.POST("/api/v1/users", mw(a, abc))

	r.Listen(":8888")
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8888", nil)
}

func home(c *router.Control) {
	fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Request.URL.Path)
}

func mw(a int, h router.Handle) router.Handle {
	return h
}

func abc(c *router.Control) {
	fmt.Fprintf(c.Writer, "a = ?")
}

func pampam(c *router.Control) {
	fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Request.URL.Path)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
