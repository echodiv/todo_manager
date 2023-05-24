package entity

// Task represent todo task in domain scope.
type Task struct {
	ID          int
	UserID      int
	IsComplete  bool
	Title       string
	Description string
}

// Storage is an interface for access to the Task data.
type Storage interface {
	Create(userID int, title string, description string) Task
	Complete(ID int) (Task, error)
	GetByUserId(userID int) []Task
	GetById(ID int) (Task, error)
}
