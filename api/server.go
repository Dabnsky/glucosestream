package api

import (
	"glucosestream/device"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Readings []device.Reading
}

func NewServer() *Server {
	return &Server{
		Readings: make([]device.Reading, 0),
	}
}

func (s *Server) Run() {
	r := gin.Default()

	r.GET("/readings", func(c *gin.Context) {
		c.JSON(http.StatusOK, s.Readings)
	})

	r.Run(":4656")
}
