package main

import (
	"fmt"

	"github.com/Ryne777/Golang_first_project/internal/config"
	"github.com/Ryne777/Golang_first_project/internal/routers"
)

func main() {
	cfg := config.GetConfig()

	r := routers.SetupRouter(cfg)
	r.Run(fmt.Sprintf(":%s", cfg.Listen.Port))

}
