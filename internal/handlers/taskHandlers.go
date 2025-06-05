package handlers

import (
	"Tasks/internal/taskService"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TaskHandler struct {
	service taskService.TaskService
}

func NewtaskHandlers(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetHandler(c echo.Context) error {
	tasks, err := h.service.GetAllTask()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not get tasks"})
	}
	return c.JSON(http.StatusOK, &tasks)
}

func (h *TaskHandler) PostHandler(c echo.Context) error {
	var tasks taskService.Task
	if err := c.Bind(&tasks); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	createdTask, err := h.service.CreateTask(tasks)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "could not create task"})
	}

	return c.JSON(http.StatusCreated, createdTask)
}

func (h *TaskHandler) PatchHandler(c echo.Context) error {
	id := c.Param("id")

	var req taskService.Task
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updatedTask, err := h.service.UpdateTask(id, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "could not update task"})
	}

	return c.JSON(http.StatusOK, &updatedTask)
}

func (h *TaskHandler) DeleteHandler(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteTask(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete task"})
	}
	return c.NoContent(http.StatusNoContent)
}
