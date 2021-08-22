package main

import (
	"github.com/makkesann/test2/app/routes"
)

func main() {
	r := routes.GetRouter()
	r.Run(":8080")
}
