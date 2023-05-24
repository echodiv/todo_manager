package storage

import (
	"fmt"
	"sync"

	"github.com/echodiv/todo_manager/internal/domain/entity"
)

type memoryStorage struct {
	todos []entity.Task
	s     sync.Mutex
}

func NewmemoryStorage() memoryStorage {
	return memoryStorage{}
}

func (m *memoryStorage) Create(userID int, title string, description string) entity.Task {
	m.s.Lock()

	newTask := entity.Task{
		ID:          len(m.todos) + 1,
		UserID:      userID,
		IsComplete:  false,
		Title:       title,
		Description: description,
	}
	m.todos = append(m.todos, newTask)

	m.s.Unlock()

	return newTask
}

func (m *memoryStorage) Complete(ID int) (entity.Task, error) {
	m.s.Lock()

	t, err := m.GetByID(ID)
	if err != nil {
		return entity.Task{}, fmt.Errorf("failed to get task: %w", err)
	}

	t.IsComplete = true

	m.s.Unlock()

	return t, nil
}

func (m *memoryStorage) GetByUserID(userID int) []entity.Task {
	var tasks []entity.Task

	for _, t := range m.todos {
		if t.UserID == userID {
			tasks = append(tasks, t)
		}
	}

	return tasks
}

func (m *memoryStorage) GetByID(ID int) (entity.Task, error) {
	for _, task := range m.todos {
		if task.ID == ID {
			return task, nil
		}
	}

	return entity.Task{}, fmt.Errorf("task not found")
}
