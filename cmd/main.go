package main

import (
	"github.com/Ryne777/Golang_first_project/internal/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}
