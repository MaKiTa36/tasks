package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

var task string

type Task struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type Model struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}

func GetHandler(c echo.Context) error {
	var tasks []Task
	if err := DB.Find(&tasks).Error; err != nil {
		return c.JSON(http.StatusBadRequest, &Task{
			Task:   "could not find the tasks",
			IsDone: false,
		})
	}
	return c.JSON(http.StatusOK, &tasks)
}

func PostHandler(c echo.Context) error {
	var tasks Task
	if err := c.Bind(&tasks); err != nil {
		return c.JSON(http.StatusBadRequest, &Task{
			Task:   "could not add the tasks",
			IsDone: false,
		})
	}
	if err := DB.Create(&tasks).Error; err != nil {
		return c.JSON(http.StatusBadRequest, &Task{
			Task:   "could not create the tasks",
			IsDone: false,
		})
	}
	return c.JSON(http.StatusOK, tasks)
}

func PatchHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Task{
			Task:   "Bad ID",
			IsDone: false,
		})
	}
	var existingTask Task
	if err := DB.First(&existingTask, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, &Task{
			Task:   "could not find the task",
			IsDone: false,
		})
	}
	var updateTask Task
	if err := c.Bind(&updateTask); err != nil {
		return c.JSON(http.StatusBadRequest, &Task{
			Task:   "Invalid input",
			IsDone: false,
		})
	}
	existingTask.Task = updateTask.Task
	existingTask.IsDone = updateTask.IsDone

	if err := DB.Save(&existingTask).Error; err != nil {
		return c.JSON(http.StatusBadRequest, &Task{
			Task:   "could not update the tasks1",
			IsDone: false,
		})
	}
	return c.JSON(http.StatusOK, &Task{
		Task:   "success",
		IsDone: true,
	})
}
func DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &Task{
			Task:   "Bad ID",
			IsDone: false,
		})
	}
	if err := DB.Delete(&Task{}, id).Error; err != nil {
		return c.JSON(http.StatusBadRequest, &Task{
			Task:   "could not delete the task",
			IsDone: false,
		})
	}
	return c.JSON(http.StatusOK, &Task{
		Task:   "success",
		IsDone: true,
	})
}

func main() {
	InitDB()
	DB.AutoMigrate(&Task{})
	e := echo.New()
	e.GET("/tasks", GetHandler)
	e.POST("/tasks", PostHandler)
	e.PATCH("/tasks/:id", PatchHandler)
	e.DELETE("/tasks/:id", DeleteHandler)
	e.Start(":8080")
}
