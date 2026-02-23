package Usecases

import (
	"task-manager/Domain"
	"task-manager/Repositories"
)

type TaskUsecase interface {
	CreateTask(title, description string) (*Domain.Task, error)
	UpdateTask(task *Domain.Task, title, description, status string) (*Domain.Task, error)
	DeleteTask(id string) error
	GetAllTasks() ([]*Domain.Task, error)
	GetTaskByID(id string) (*Domain.Task, error)
}

type taskUsecase struct {
	repo Repositories.TaskRepository
}

func NewTaskUsecase(r Repositories.TaskRepository) TaskUsecase {
	return &taskUsecase{repo: r}
}

func (u *taskUsecase) CreateTask(title, description string) (*Domain.Task, error) {
	task := Domain.NewTask("", title, description)
	return u.repo.Create(task)
}

func (u *taskUsecase) UpdateTask(t *Domain.Task, title, description, status string) (*Domain.Task, error) {
	t.Update(title, description, status)
	return u.repo.Update(t)
}

func (u *taskUsecase) DeleteTask(id string) error {
	return u.repo.Delete(id)
}

func (u *taskUsecase) GetAllTasks() ([]*Domain.Task, error) {
	return u.repo.GetAll()
}

func (u *taskUsecase) GetTaskByID(id string) (*Domain.Task, error) {
	return u.repo.GetByID(id)
}