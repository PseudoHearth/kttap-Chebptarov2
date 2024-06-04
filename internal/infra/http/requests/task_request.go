package requests

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type TaskRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description"`
	Deadline    *int64  `json:"deadline,omitempty"`
}

func (r TaskRequest) ToDomainModel() (interface{}, error) {
	var deadline *time.Time
	if r.Deadline != nil {
		dl := time.Unix(*r.Deadline, 0)
		deadline = &dl
	}

	return domain.Task{
		Title:       r.Title,
		Description: r.Description,
		Deadline:    deadline,
	}, nil
}

type TaskUpdateRequest struct {
	Id          uint64  `json:"id" validate:"required"`
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description"`
	Status      string  `json:"status" validate:"required"`
	Deadline    *int64  `json:"deadline,omitempty"`
}

func (r TaskUpdateRequest) ToDomainModel() (interface{}, error) {
	var deadline *time.Time
	if r.Deadline != nil {
		dl := time.Unix(*r.Deadline, 0)
		deadline = &dl
	}

	return domain.Task{
		Id:          r.Id,
		Title:       r.Title,
		Description: r.Description,
		Status:      domain.TaskStatus(r.Status),
		Deadline:    deadline,
	}, nil
}
