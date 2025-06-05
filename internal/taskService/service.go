package taskService

import (
	"errors"
	"log"
	"time"
)

type TaskService interface {
	CreateTask(task Task) (*Task, error)
	GetAllTask() ([]Task, error)
	GetTaskById(id string) (Task, error)
	UpdateTask(id string, updates Task) (Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo TaskRepository
}

var ErrEmptyTask = errors.New("task text cannot be empty")

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(task Task) (*Task, error) {
	log.Printf("Creating task: %+v", task)
	if task.Task == "" {
		return nil, ErrEmptyTask
	}
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	if err := s.repo.CreateTask(task); err != nil {
		log.Printf("Error creating task: %v", err)
		return nil, err
	}
	return &task, nil
}

func (s *taskService) GetAllTask() ([]Task, error) {
	return s.repo.GetAllTasks()
}

func (s *taskService) GetTaskById(id string) (Task, error) {
	return s.repo.GetTaskById(id)
}

func (s *taskService) UpdateTask(id string, updates Task) (Task, error) {
	task, err := s.repo.GetTaskById(id)
	if err != nil {
		return Task{}, err
	}

	if updates.Task != "" {
		task.Task = updates.Task
	}
	task.IsDone = updates.IsDone

	err = s.repo.UpdateTask(task)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *taskService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
