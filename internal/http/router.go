package http

import (
	"log"
	"os"

	_ "todo-list/docs" // Import generated docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	router.Use(CORSMiddleware())

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Error run router: %s", err)
		return
	}
}
