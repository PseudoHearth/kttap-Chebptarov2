package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type TaskService interface {
	Save(t domain.Task) (domain.Task, error)
	FindByUserId(uId uint64) ([]domain.Task, error)
	FindById(id uint64) (domain.Task, error)
	Update(t domain.Task) (domain.Task, error)
	Delete(taskId uint64) error
}

type taskService struct {
	taskRepo database.TaskRepository
}

func (s taskService) Delete(taskId uint64) error {
	return s.taskRepo.Delete(taskId)
}

func NewTaskService(tr database.TaskRepository) TaskService {
	return taskService{
		taskRepo: tr,
	}
}

func (s taskService) Save(t domain.Task) (domain.Task, error) {
	task, err := s.taskRepo.Save(t)
	if err != nil {
		log.Printf("TaskService -> Save: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}

func (s taskService) FindByUserId(uId uint64) ([]domain.Task, error) {
	tasks, err := s.taskRepo.FindByUserId(uId)
	if err != nil {
		log.Printf("TaskService -> FindByUserId: %s", err)
		return nil, err
	}
	return tasks, nil
}

func (s taskService) FindById(id uint64) (domain.Task, error) {
	task, err := s.taskRepo.FindById(id)
	if err != nil {
		log.Printf("TaskService -> FindById: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}

func (s taskService) Update(t domain.Task) (domain.Task, error) {
	task, err := s.taskRepo.Update(t)
	if err != nil {
		log.Printf("TaskService -> Update: %s", err)
		return domain.Task{}, err
	}
	return task, nil
}
