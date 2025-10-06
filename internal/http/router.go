package http

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Handler *Handler
}

func NewRoute() *Route {
	return &Route{
		Handler: newHandler(),
	}
}

func (r *Route) RouteRun() {

	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "release"
	}
	gin.SetMode(ginMode)

	router := gin.Default()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	//router.Use(CORSMiddleware())

	router.POST("/register", r.Handler.RegisterUser)
	router.POST("/login", r.Handler.LoginUser)

	todos := router.Group("todos")
	todos.Use(AuthMiddleware())
	todos.GET("/", r.Handler.ListTodos)
	todos.POST("/create", r.Handler.CreateTodo)
	todos.PUT("/:id", r.Handler.UpdateTodo)
	todos.DELETE("/:id", r.Handler.DeleteTodo)

	task := router.Group("task")
	task.Use(AuthMiddleware())
	task.PUT("/:todoId/:taskId", r.Handler.UpdateStatusTask)
	task.POST("/:id", r.Handler.CreateTask)

	log.Printf("Starting server on port %s", ":"+port)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Error run router: %s", err)
		return
	}
}
