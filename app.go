package main

import "github.com/codehell/expenses/routes"

func main() {
	/*r := routes.Routes()
	http.ListenAndServe(":3000", r)*/
	r := routes.Router()
	_ = r.Run(":3000")
}
