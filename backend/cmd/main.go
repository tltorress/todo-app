package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tltorress/todo-app/internal/database"
	"github.com/tltorress/todo-app/internal/todoapp/handlers"
	"github.com/tltorress/todo-app/internal/todoapp/task"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	app := gin.New()

	cfg, err := NewConfig()
	if err != nil {
		return err
	}

	db, err := database.NewDatabase("root", cfg.DB.Pass, cfg.DB.Host, cfg.DB.Port, cfg.DB.Database, "mysql")
	if err != nil {
		return err
	}

	app.GET("/ping", handlers.Ping())

	app.GET("/tasks", handlers.GetTasksHandler(task.NewTaskRepository(db)))

	app.GET("/tasks/:id", handlers.GetTaskHandler(task.NewTaskRepository(db)))

	app.POST("/tasks", handlers.CreateTaskHandler(task.NewTaskRepository(db)))

	app.PUT("/tasks/:id", handlers.UpdateTaskHandler(task.NewTaskRepository(db)))

	app.DELETE("/tasks/:id", handlers.DeleteTasksHandler(task.NewTaskRepository(db)))

	err = app.Run("localhost:9090")
	if err != nil {
		return err
	}

	return nil
}
