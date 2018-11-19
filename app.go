package main

import (
	"github.com/codehell/expenses/routes"
	"net/http"
)

func main() {
	r := routes.Routes()
	http.ListenAndServe(":3000", r)
}
