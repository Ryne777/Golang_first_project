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
	db.AutoMigrate(&model.ToDo{}, &model.Author{})
	cnt := controller.NewToDoController(db)
	cnt2 := controller.NewAuthorController(db)
	r := routers.SetupRouter(cfg, cnt, cnt2)
	r.Run(fmt.Sprintf(":%s", cfg.Listen.Port))

}
