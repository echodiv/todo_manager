package rest

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/echodiv/todo_manager/internal/domain/usecase"
)

type CreateRequest struct {
	UserId      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CompleteRequest struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
}

type TaskResponse struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	IsComplete  bool   `json:"complete"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Service interface {
	Create(taskRequest usecase.TaskCreateRequest) usecase.Task
	Complete(userId, Id int) (usecase.Task, error)
	GetForUser(userId int) []usecase.Task
}

type Task struct {
	in Service
}

func NewTaskService(in usecase.TaskInteractor) Task {
	return Task{
		in: in,
	}
}

func (t Task) Create(c echo.Context) error {
	requestBody := new(CreateRequest)

	if err := c.Bind(requestBody); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	todo := t.in.Create(createRequestToInteractorsRequest(requestBody))

	return c.JSON(http.StatusCreated, interactorsTaskToTaskResponse(todo))
}

func (t Task) Complete(c echo.Context) error {
	requestBody := new(CompleteRequest)

	if err := c.Bind(requestBody); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	todo, err := t.in.Complete(requestBody.UserId, requestBody.Id)

	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusCreated, interactorsTaskToTaskResponse(todo))
}

func (t Task) GetAllTaskForUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	todos := t.in.GetForUser(id)
	return c.JSON(http.StatusCreated, interactorsTasksToTaskResponses(todos))
}

func createRequestToInteractorsRequest(createRequest *CreateRequest) usecase.TaskCreateRequest {
	return usecase.TaskCreateRequest{
		UserID:      createRequest.UserId,
		Title:       createRequest.Title,
		Description: createRequest.Description,
	}
}

func interactorsTaskToTaskResponse(interactorsTask usecase.Task) TaskResponse {
	return TaskResponse{
		Id:          interactorsTask.Id,
		UserId:      interactorsTask.UserId,
		IsComplete:  interactorsTask.IsComplete,
		Title:       interactorsTask.Title,
		Description: interactorsTask.Description,
	}
}

func interactorsTasksToTaskResponses(interactorsTasks []usecase.Task) []TaskResponse {
	var tasks []TaskResponse

	for i := range interactorsTasks {
		tasks = append(tasks, interactorsTaskToTaskResponse(interactorsTasks[i]))
	}

	return tasks
}
