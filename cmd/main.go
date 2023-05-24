package main

import (
	"github.com/echodiv/todo_manager/app/domain/task/storage"
	"github.com/echodiv/todo_manager/app/interactors"
	"github.com/echodiv/todo_manager/app/services/rest"
)

func main() {
	httpServer := NewHTTPServer()
	httpServer.Start()
	for {
	}
}

func createNewService() rest.Task {
	s := storage.NewMemoryStorage()
	i := interactors.NewTaskInteractor(&s)
	return rest.NewTaskService(i)
}
