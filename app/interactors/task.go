package interactors

import (
	"fmt"

	"github.com/echodiv/todo/app/domain/task"
)

type Task struct {
	Id          int
	UserId      int
	IsComplete  bool
	Title       string
	Description string
}

type TaskCreateRequest struct {
	UserID      int
	Title       string
	Description string
}

type TaskInteractor struct {
	m task.Storage
}

func NewTaskInteractor(storage task.Storage) TaskInteractor {
	return TaskInteractor{
		m: storage,
	}
}

func (t TaskInteractor) Create(taskRequest TaskCreateRequest) Task {
	domainTask := t.m.Create(taskRequest.UserID, taskRequest.Title, taskRequest.Description)

	return domainTaskToTask(domainTask)
}

func (t TaskInteractor) Complete(userId, Id int) (Task, error) {
	err := t.checkAccessAllowed(userId, Id)
	if err != nil {
		return Task{}, fmt.Errorf("failed access: %w", err)
	}

	domainTask, err := t.m.Complete(Id)
	if err != nil {
		return Task{}, fmt.Errorf("failed to complete task: %w", err)
	}

	return domainTaskToTask(domainTask), nil
}

func (t TaskInteractor) GetForUser(userId int) []Task {
	localTasks := t.m.GetByUserId(userId)

	return domainTasksToTasks(localTasks)
}

func (t TaskInteractor) checkAccessAllowed(userID int, Id int) error {
	domainTask, err := t.m.GetById(Id)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}
	localTask := domainTaskToTask(domainTask)

	if localTask.UserId != userID {
		return fmt.Errorf("premission denied")
	}

	return nil
}

func domainTaskToTask(domainTask task.Task) Task {
	return Task{
		Id:          domainTask.Id,
		UserId:      domainTask.UserId,
		IsComplete:  domainTask.IsComplete,
		Title:       domainTask.Title,
		Description: domainTask.Description,
	}
}

func domainTasksToTasks(domainTasks []task.Task) (tasks []Task) {
	for ti := range domainTasks {
		tasks = append(tasks, domainTaskToTask(domainTasks[ti]))
	}

	return tasks
}
