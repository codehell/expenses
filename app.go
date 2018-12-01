package main

import "github.com/codehell/expenses/routes"

func main() {
	r := routes.Router()
	_ = r.Run(":3000")
}
