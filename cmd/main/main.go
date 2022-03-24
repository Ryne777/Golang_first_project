package main

import (
	"fmt"

	"github.com/Ryne777/Golang_first_project/internal/config"
	"github.com/Ryne777/Golang_first_project/internal/controller"
	"github.com/Ryne777/Golang_first_project/internal/model"
	"github.com/Ryne777/Golang_first_project/internal/routers"
	"github.com/Ryne777/Golang_first_project/pkg/bdClient"
)

func main() {
	cfg := config.GetConfig()
	db := bdClient.GetConnectionDb(cfg)
	db.AutoMigrate(&model.ToDo{})
	cnt := controller.NewToDoController(db)
	r := routers.SetupRouter(cfg, cnt)
	r.Run(fmt.Sprintf(":%s", cfg.Listen.Port))

}
