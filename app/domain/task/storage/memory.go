package storage

import (
	"fmt"
	"sync"

	"github.com/echodiv/todo/app/domain/task"
)

type MemoryStorage struct {
	todos []task.Task
	s     sync.Mutex
}

func NewMemoryStorage() MemoryStorage {
	return MemoryStorage{}
}

func (m *MemoryStorage) Create(userID int, title string, description string) task.Task {
	m.s.Lock()

	newTask := task.Task{
		Id:          len(m.todos) + 1,
		UserId:      userID,
		IsComplete:  false,
		Title:       title,
		Description: description,
	}
	m.todos = append(m.todos, newTask)

	m.s.Unlock()

	return newTask
}

func (m *MemoryStorage) Complete(ID int) (task.Task, error) {
	m.s.Lock()

	t, err := m.GetById(ID)
	if err != nil {
		return task.Task{}, fmt.Errorf("failed to get task: %w", err)
	}

	t.IsComplete = true

	m.s.Unlock()

	return t, nil
}

func (m *MemoryStorage) GetByUserId(userID int) []task.Task {
	var tasks []task.Task

	for _, t := range m.todos {
		if t.UserId == userID {
			tasks = append(tasks, t)
		}
	}

	return tasks
}

func (m *MemoryStorage) GetById(ID int) (task.Task, error) {
	for _, task := range m.todos {
		if task.Id == ID {
			return task, nil
		}
	}

	return task.Task{}, fmt.Errorf("task not found")
}
