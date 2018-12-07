package main

import "expenses/routes"

func main() {
	r := routes.Router()
	_ = r.Run(":3000")
}
