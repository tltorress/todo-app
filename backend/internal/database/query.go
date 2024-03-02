package database

const (
	GetTasksQuery = "SELECT id, title, description, finish_date, created_date, last_updated FROM tasks"

	GetTaskQuery = "SELECT id, title, description, finish_date, created_date, last_updated FROM tasks WHERE id = ?"

	DeleteTasksQuery = "DELETE FROM tasks WHERE id = ?"

	UpdateTaskQuery = `UPDATE tasks
	SET title = ?, 
	description = ?,
	finish_date = ?
	WHERE id = ?; `

	CreateTaskQuery = `INSERT INTO tasks(title, description)
	VALUES(?,?)`
)
