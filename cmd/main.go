package main

import (
	"Tasks/internal/db"
	"Tasks/internal/handlers"
	"Tasks/internal/taskService"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log")

func main() {
		db.InitDB()
		db.DB.AutoMigrate(&taskService.Task{})

		repo := taskService.NewTaskRepository(db.DB)
		service := taskService.NewTaskService(repo)

		handler := handlers.NewtaskHandlers(service)

		// Инициализируем echo
		e := echo.New()

		// используем Logger и Recover
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		// Прикол для работы в echo. Передаем и регистрируем хендлер в echo
		strictHandler := tasks.NewStrictHandler(handler, nil) // тут будет ошибка
		tasks.RegisterHandlers(e, strictHandler)

		if err := e.Start(":8080"); err != nil {
			log.Fatalf("failed to start with err: %v", err)
		}
	}

	e.GET("/tasks", taskHandlers.GetHandler)
	e.POST("/tasks", taskHandlers.PostHandler)
	e.PATCH("/tasks/:id", taskHandlers.PatchHandler)
	e.DELETE("/tasks/:id", taskHandlers.DeleteHandler)
	e.Start(":8080")
}

