package controllers

import (
	"net/http"

	"github.com/zaahidali/task_manager_api/data"
	"github.com/zaahidali/task_manager_api/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "tasks": data.Tasks})
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	for _, task := range data.Tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range data.Tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				data.Tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				data.Tasks[i].Description = updatedTask.Description
			}
			if !updatedTask.DueDate.IsZero() {
				data.Tasks[i].DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != "" {
				data.Tasks[i].Status = updatedTask.Status
			}
			c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	for i, task := range data.Tasks {
		if task.ID == id {
			data.Tasks = append(data.Tasks[:i], data.Tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.Tasks = append(data.Tasks, newTask)
	c.JSON(http.StatusCreated, gin.H{"message": "Task created", "task": newTask})
}
