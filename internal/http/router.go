package http

import (
	"net/http"
	"todo-list/internal/config"
	_ "todo-list/docs"
	"todo-list/pkg"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Route struct {
	Handler *Handler
	cfg     *config.Config
}

func NewRoute(cfg *config.Config) *Route {
	return &Route{
		Handler: newHandler(cfg),
		cfg:     cfg,
	}
}

func (r *Route) RouteRun() {
	gin.SetMode(r.cfg.App.GinMode)

	router := gin.New()

	router.Use(RecoveryMiddleware())
	router.Use(CORSMiddleware(r.cfg.App.CORSOrigins))
	router.Use(RequestLoggerMiddleware())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

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
		Str("port", r.cfg.App.Port).
		Str("mode", r.cfg.App.GinMode).
		Msg("Starting server")

	if err := router.Run(":" + r.cfg.App.Port); err != nil {
		pkg.Logger.Fatal().
			Err(err).
			Msg("Failed to start server")
	}
}
