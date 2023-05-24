package servicemanager

// Serive is interface for management.
type Service interface {
	Start()
	Stop()
	Restart()
}

// Manager - service manager.
type Manager struct {
	Services []Service
}

func NewManager() *Manager {
	return &Manager{}
}

func (s *Manager) Run() {

}
