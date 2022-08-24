package httpServer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"AntonyChR/golang-hexagonal-arch-template/internal/core/application"
	"AntonyChR/golang-hexagonal-arch-template/internal/core/domain"

	"github.com/gin-gonic/gin"
)

type Server struct {
	todoService *application.TodoService
	port        string
}

func (s *Server) Start() {
	server := gin.Default()

	server.GET("/id/:id", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		data := s.todoService.GetById(c.Param("id"))
		c.JSON(http.StatusOK, data)
	})

	server.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		data, err := s.todoService.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error getting task",
			})
		} else {
			c.JSON(http.StatusOK, data)
		}
	})

	server.PATCH("/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := s.todoService.MarkAsDone(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error updating task",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Task updated",
			})
		}

	})

	server.POST("/", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		var test domain.Todo
		er := json.Unmarshal(body, &test)
		if er != nil {
			fmt.Println(er)
		}
		id, err := s.todoService.Create(test)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error creating task",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	server.DELETE("/:id", func(c *gin.Context) {
		id := c.Param("id")
		err := s.todoService.DeleteById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Error deleting task",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Deleted task",
		})
	})
	server.Run(s.port)
}

func Initialize(port string, service *application.TodoService) *Server {
	return &Server{
		todoService: service,
		port:        port,
	}
}
