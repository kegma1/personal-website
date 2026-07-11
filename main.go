package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"personal_website/views"
)

func main() {
	component := views.Hello("John")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
