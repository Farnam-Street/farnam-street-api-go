package main

import (
	Routers "farnam-street-api-go/routers"
	"fmt"
)

func main() {
	fmt.Println("Sample GO Lang Program")
	r := Routers.SetupRouter()

	// type Job interface {
	// 	Run()
	// }
	r.Run(":8080")
}
