package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/zaahidali/task_manager_api/data"
	"github.com/zaahidali/task_manager_api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	tasks, err := data.GetAllTasks(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "tasks": tasks})
}

func GetTaskById(c *gin.Context) {
	id := c.Param("id")
	task, err := data.GetTaskById(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := data.UpdateTask(context.Background(), id, updatedTask)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or update failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task Updated"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := data.DeleteTask(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = primitive.NewObjectID()
	if newTask.DueDate.IsZero() {
		newTask.DueDate = time.Now()
	}

	_, err := data.CreateTask(context.Background(), newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Task created", "task": newTask})
}
