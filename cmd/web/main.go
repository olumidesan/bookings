package main

import (
	"fmt"
	"net/http"

	"github.com/olumidesan/bookings/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting HTTP Server...")
	_ = http.ListenAndServe(":3000", nil)
}
