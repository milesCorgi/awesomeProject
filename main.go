package main

import (
	"awesomeProject/routes"
)

func main() {
	r := routes.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
