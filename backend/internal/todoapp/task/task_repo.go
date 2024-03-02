package task

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/tltorress/todo-app/internal/database"
)

var (
	ErrTaskNotFound  = errors.New("task not found")
	ErrInvalidTaskID = errors.New("task id must be a valid integer")
	ErrCreateTask    = errors.New("cant create task")
	ErrUpdateTask    = errors.New("cant update task")
	ErrDeleteTask    = errors.New("cant delete tasks")
)

type Repository interface {
	UpdateTask(task Task, id uint64) error
	GetTask(id uint64) (Task, error)
	DeleteTasks(id uint64) error
	CreateTask(task Task) error
	GetTasks() ([]Task, error)
}

func NewTaskRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *sql.DB
}

func (r *repository) CreateTask(task Task) error {
	result, err := r.db.Exec(database.CreateTaskQuery, task.Title, task.Description)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	if id == 0 {
		return ErrCreateTask
	}

	return nil
}

func (r *repository) UpdateTask(task Task, id uint64) error {

	result, err := r.db.Exec(database.UpdateTaskQuery, task.Title, task.Description, task.FinishDate, task.ID)
	if err != nil {
		return err
	}

	q, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if q <= 0 {
		return ErrUpdateTask
	}

	return nil
}

func (r *repository) GetTask(id uint64) (Task, error) {

	rows, err := r.db.Query(database.GetTaskQuery, id)
	if err != nil {
		return Task{}, err
	}

	var t Task
	for rows.Next() {
		_ = rows.Scan(&t.ID, &t.Title, &t.Description, &t.FinishDate, &t.LastUpdated, &t.CreatedDate)
	}

	return t, nil
}

func (r *repository) GetTasks() ([]Task, error) {
	rows, err := r.db.Query(database.GetTasksQuery)
	if err != nil {
		return nil, err
	}

	var tks []Task
	for rows.Next() {
		t := Task{}
		err = rows.Scan(&t.ID, &t.Title, &t.Description, &t.FinishDate, &t.CreatedDate, &t.LastUpdated)
		if err != nil {
			fmt.Println(err)
		}

		tks = append(tks, t)
	}

	fmt.Println(tks)

	return tks, nil
}

func (r *repository) DeleteTasks(id uint64) error {

	result, err := r.db.Exec(database.DeleteTasksQuery, id)
	if err != nil {
		return err
	}

	q, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if q <= 0 {
		return ErrDeleteTask
	}

	return nil
}

func (r *repository) buildDeleteTasksQuery(ids []uint64) (string, any) {

	query := database.DeleteTasksQuery
	var args = make([]any, len(ids))
	for i := range ids {
		args[i] = ids[i]
		print(query)
		if i < len(ids)-1 {
			query += "?,"
			continue
		}
		query += "?)"
	}

	fmt.Println(query)

	return query, args
}
