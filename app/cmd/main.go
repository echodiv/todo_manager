package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/echodiv/todo/app/domain/task/storage"
	"github.com/echodiv/todo/app/interactors"
	"github.com/echodiv/todo/app/services/rest"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	buildEndpoints(e)

	e.Start(":8888")
}

func createNewService() rest.Task {
	s := storage.NewMemoryStorage()
	i := interactors.NewTaskInteractor(&s)
	return rest.NewTaskService(i)
}

func buildEndpoints(e *echo.Echo) {
	r := createNewService()
	e.POST("/create", r.Create)
	e.POST("/update", r.Complete)
	e.GET("/tasks/:id", r.GetAllTaskForUser)
}
