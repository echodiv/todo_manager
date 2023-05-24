package usecase

import (
	"fmt"

	"github.com/echodiv/todo_manager/internal/domain/entity"
)

// Storage is an interface for access to the Task data.
type Storage interface {
	Create(userID int, title string, description string) entity.Task
	Complete(ID int) (entity.Task, error)
	GetByUserID(userID int) []entity.Task
	GetByID(ID int) (entity.Task, error)
}

type TaskCreateRequest struct {
	UserID      int
	Title       string
	Description string
}

type TaskInteractor struct {
	m entity.Storage
}

func NewTaskInteractor(storage entity.Storage) TaskInteractor {
	return TaskInteractor{
		m: storage,
	}
}

func (t TaskInteractor) Create(taskRequest TaskCreateRequest) entity.Task {

	return t.m.Create(taskRequest.UserID, taskRequest.Title, taskRequest.Description)
}

func (t TaskInteractor) Complete(userID, ID int) (entity.Task, error) {
	err := t.checkAccessAllowed(userID, ID)
	if err != nil {
		return entity.Task{}, fmt.Errorf("failed access: %w", err)
	}

	task, err := t.m.Complete(ID)
	if err != nil {
		return entity.Task{}, fmt.Errorf("failed to complete task: %w", err)
	}

	return task, nil
}

func (t TaskInteractor) GetForUser(userID int) []entity.Task {

	return t.m.GetByUserID(userID)
}

func (t TaskInteractor) checkAccessAllowed(userID int, ID int) error {
	task, err := t.m.GetByID(ID)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	if task.UserID != userID {
		return fmt.Errorf("premission denied")
	}

	return nil
}
