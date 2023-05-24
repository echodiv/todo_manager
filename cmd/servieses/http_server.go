package servieses

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// HTTPServer - WEB server for task management.
type HTTPServer struct {
}

// NewHTTPServer - create new HTTP server for task managenet.
func NewHTTPServer() {
	e := echo.New()

	e.Use(middleware.Logger())

	setupRoutes(e)

	e.Start(":8888")
}

func setupRoutes(e *echo.Echo) {
	r := createNewService()
	e.POST("/create", r.Create)
	e.POST("/update", r.Complete)
	e.GET("/tasks/:id", r.GetAllTaskForUser)
}

// Start - run HTTP server in backgroud.
func (s *HTTPServer) Start() {

}

// Stop - stop HTTP server.
func (s *HTTPServer) Stop() {

}

// Restart - stop HTTP server and start it again.
func (s *HTTPServer) Restart() {

}
