package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func MainRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello world"})
}

func GetAllTasks(ctx *gin.Context) {
	rows, err := DB.Query("SELECT id, title FROM tasks")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task

		if err := rows.Scan(&task.Id, &task.Title); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	ctx.IndentedJSON(http.StatusOK, tasks)
}

func CreateTask(ctx *gin.Context) {
	var newTask Task

	if err := ctx.BindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := DB.Exec("INSERT INTO tasks (title) VALUES (?)", newTask.Title)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id, err := result.LastInsertId()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newTask.Id = int(id)
	ctx.JSON(http.StatusCreated, newTask)
}

func GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	var task Task
	row := DB.QueryRow("SELECT id, title FROM tasks WHERE id = ?", id)

	if err := row.Scan(&task.Id, &task.Title); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func DeleteTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := DB.Exec("DELETE FROM tasks WHERE id = ?", id)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func UpdateTaskById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var updatedTask Task

	if err := ctx.BindJSON(&updatedTask); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	_, err := DB.Exec("UPDATE tasks SET title = ? WHERE id = ?", updatedTask.Title, id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask.Id = id
	ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
}
