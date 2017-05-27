package main

import (
	"fmt"

	"github.com/takama/router"
)

func home(c *router.Control) {
	fmt.Fprintf(c.Writer, "URL.Path = %q\n", c.Request.URL.Path)
}
