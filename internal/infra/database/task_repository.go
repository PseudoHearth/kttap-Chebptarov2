package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const TasksTableName = "tasks"

type task struct {
	Id          uint64     `db:"id,omitempty"`
	UserId      uint64     `db:"user_id"`
	Title       string     `db:"title"`
	Description *string    `db:"description,omitempty"`
	Status      string     `db:"status"`
	Deadline    *time.Time `db:"deadline,omitempty"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date,omitempty"`
}

type TaskRepository interface {
	Save(t domain.Task) (domain.Task, error)
	FindByUserId(uId uint64) ([]domain.Task, error)
	FindById(id uint64) (domain.Task, error)
	Update(t domain.Task) (domain.Task, error)
	Delete(id uint64) error
}

type taskRepository struct {
	sess db.Session
	coll db.Collection
}

func NewTaskRepository(session db.Session) TaskRepository {
	return taskRepository{
		sess: session,
		coll: session.Collection(TasksTableName),
	}
}

func (r taskRepository) Save(t domain.Task) (domain.Task, error) {
	task := r.domainToDb(t)
	if task.Id == 0 {
		err := r.coll.InsertReturning(&task)
		if err != nil {
			return domain.Task{}, err
		}
	} else {
		err := r.coll.UpdateReturning(&task)
		if err != nil {
			return domain.Task{}, err
		}
	}
	return r.dbToDomain(task), nil
}

func (r taskRepository) FindByUserId(uId uint64) ([]domain.Task, error) {
	var tasks []task
	err := r.coll.Find(db.Cond{"user_id": uId}).All(&tasks)
	if err != nil {
		return nil, err
	}
	var domainTasks []domain.Task
	for _, t := range tasks {
		domainTasks = append(domainTasks, r.dbToDomain(t))
	}
	return domainTasks, nil
}

func (r taskRepository) FindById(id uint64) (domain.Task, error) {
	var task task
	err := r.coll.Find(db.Cond{"id": id}).One(&task)
	if err != nil {
		return domain.Task{}, err
	}
	return r.dbToDomain(task), nil
}

func (r taskRepository) Update(t domain.Task) (domain.Task, error) {
	tsk := r.domainToDb(t)
	tsk.UpdatedDate = time.Now()
	err := r.coll.Find(tsk.Id).Update(tsk)
	if err != nil {
		return domain.Task{}, err
	}
	return r.dbToDomain(tsk), nil
}

func (r taskRepository) Delete(id uint64) error {
	err := r.coll.Find(db.Cond{"id": id}).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (r taskRepository) domainToDb(t domain.Task) task {
	return task{
		Id:          t.Id,
		UserId:      t.UserId,
		Title:       t.Title,
		Description: t.Description,
		Status:      string(t.Status),
		Deadline:    t.Deadline,
		CreatedDate: t.CreatedDate,
		UpdatedDate: t.UpdatedDate,
		DeletedDate: t.DeletedDate,
	}
}

func (r taskRepository) dbToDomain(t task) domain.Task {
	return domain.Task{
		Id:          t.Id,
		UserId:      t.UserId,
		Title:       t.Title,
		Description: t.Description,
		Status:      domain.TaskStatus(t.Status),
		Deadline:    t.Deadline,
		CreatedDate: t.CreatedDate,
		UpdatedDate: t.UpdatedDate,
		DeletedDate: t.DeletedDate,
	}
}
