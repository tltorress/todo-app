package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tltorress/todo-app/internal/todoapp/task"
)

func GetTasksHandler(repo task.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		tasks, err := repo.GetTasks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			fmt.Println(err)
			return
		}
		c.JSON(http.StatusOK, tasks)
	}
}

func GetTaskHandler(repo task.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": task.ErrInvalidTaskID,
			})
			return
		}

		t, err := repo.GetTask(id)
		if err != nil {
			switch err {
			case task.ErrTaskNotFound:
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				fmt.Println(err)
			}
			return
		}

		c.JSON(http.StatusOK, t)

	}
}

func CreateTaskHandler(repo task.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := task.Task{}
		err := c.BindJSON(&t)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": task.ErrInvalidTaskID,
			})
			return
		}

		err = repo.CreateTask(t)
		if err != nil {
			switch err {
			default:
				c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
				fmt.Println(err)
			}
			return
		}

		c.JSON(http.StatusCreated, "OK")
	}
}

func UpdateTaskHandler(repo task.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": task.ErrInvalidTaskID,
			})
			return
		}

		fmt.Println(id)

		completeTask, err := repo.GetTask(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": task.ErrTaskNotFound,
			})
			return
		}

		err = c.BindJSON(&completeTask)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = repo.UpdateTask(completeTask, id)
		if err != nil {
			switch err {
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				fmt.Println(err)
				return
			}
		}

		c.JSON(http.StatusOK, completeTask)
	}

}

func DeleteTasksHandler(repo task.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {

		taskID, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": task.ErrInvalidTaskID,
			})
			return
		}

		err = repo.DeleteTasks(taskID)
		if err != nil {
			switch err {
			case task.ErrTaskNotFound:
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
				fmt.Println(err)
			}
			return
		}

		c.JSON(http.StatusOK, taskID)
	}
}

type DeleteRequest struct {
	TaskIDs []uint64 `json:"task_ids"`
}
