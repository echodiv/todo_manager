package rest

import (
	"net/http"
	"strconv"

	"github.com/echodiv/todo/app/interactors"
	"github.com/labstack/echo/v4"
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
	Create(taskRequest interactors.TaskCreateRequest) interactors.Task
	Complete(userId, Id int) (interactors.Task, error)
	GetForUser(userId int) []interactors.Task
}

type Task struct {
	in Service
}

func NewTaskService(in interactors.TaskInteractor) Task {
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

func createRequestToInteractorsRequest(createRequest *CreateRequest) interactors.TaskCreateRequest {
	return interactors.TaskCreateRequest{
		UserID:      createRequest.UserId,
		Title:       createRequest.Title,
		Description: createRequest.Description,
	}
}

func interactorsTaskToTaskResponse(interactorsTask interactors.Task) TaskResponse {
	return TaskResponse{
		Id:          interactorsTask.Id,
		UserId:      interactorsTask.UserId,
		IsComplete:  interactorsTask.IsComplete,
		Title:       interactorsTask.Title,
		Description: interactorsTask.Description,
	}
}

func interactorsTasksToTaskResponses(interactorsTasks []interactors.Task) []TaskResponse {
	var tasks []TaskResponse

	for i := range interactorsTasks {
		tasks = append(tasks, interactorsTaskToTaskResponse(interactorsTasks[i]))
	}

	return tasks
}
