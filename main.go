package main

import (
	"configs/routers"
)

func main() {
	r := routers.GetRouters()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
