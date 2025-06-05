package main

import (
	"Tasks/internal/db"
	"Tasks/internal/handlers"
	"Tasks/internal/taskService"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("couldn't connect to database: %v", err)
	}
	
	e := echo.New()

	taskRepo := taskService.NewTaskRepository(database)
	taskService := taskService.NewTaskService(taskRepo)
	taskHandlers := handlers.NewtaskHandlers(taskService)

	e.GET("/tasks", taskHandlers.GetHandler)
	e.POST("/tasks", taskHandlers.PostHandler)
	e.PATCH("/tasks/:id", taskHandlers.PatchHandler)
	e.DELETE("/tasks/:id", taskHandlers.DeleteHandler)
	e.Start(":8080")
}
