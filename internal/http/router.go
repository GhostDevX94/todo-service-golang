package http

import (
	"os"

	_ "todo-list/docs"
	"todo-list/pkg"

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

	router := gin.New()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	router.Use(RecoveryMiddleware())
	router.Use(CORSMiddleware())
	router.Use(RequestLoggerMiddleware())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authGroup := router.Group("/")

	{
		authGroup.POST("/register", r.Handler.RegisterUser)
		authGroup.POST("/login", r.Handler.LoginUser)
	}

	todos := router.Group("todos")
	todos.Use(AuthMiddleware(r.Handler.Services.JWTManager))
	{
		todos.GET("/", r.Handler.ListTodos)
		todos.POST("/create", r.Handler.CreateTodo)
		todos.PUT("/:id", r.Handler.UpdateTodo)
		todos.DELETE("/:id", r.Handler.DeleteTodo)
	}

	task := router.Group("task")
	task.Use(AuthMiddleware(r.Handler.Services.JWTManager))
	{
		task.PUT("/:todoId/:taskId", r.Handler.UpdateStatusTask)
		task.POST("/:id", r.Handler.CreateTask)
	}

	pkg.Logger.Info().
		Str("port", port).
		Str("mode", ginMode).
		Msg("Starting server")

	err := router.Run(":" + port)
	if err != nil {
		pkg.Logger.Fatal().
			Err(err).
			Msg("Failed to start server")
		return
	}
}
