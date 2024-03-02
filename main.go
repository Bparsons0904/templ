package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := indexPage()

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")

	// handlerRoute := templ.Handler(component)
	// c1 := blue()
	// c2 := red()
	// route := templ.NewCSSMiddleware(handlerRoute, c1, c2)
	// http.Handle("/", route)

	http.ListenAndServe(":3000", nil)
}
