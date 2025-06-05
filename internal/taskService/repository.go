package taskService

import "gorm.io/gorm"

type TaskRepository interface {
	CreateTask(task Task) error
	GetAllTasks() ([]Task, error)
	GetTaskById(taskId string) (Task, error)
	UpdateTask(task Task) error
	DeleteTask(id string) error
}

type taskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{DB: db}
}

func (r *taskRepository) CreateTask(task Task) error {
	return r.DB.Create(&task).Error
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.DB.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTaskById(taskId string) (Task, error) {
	var task Task
	err := r.DB.First(&task, taskId).Error
	return task, err
}

func (r *taskRepository) UpdateTask(task Task) error {
	return r.DB.Save(&task).Error
}

func (r *taskRepository) DeleteTask(id string) error {
	return r.DB.Delete(&Task{}, id).Error
}
