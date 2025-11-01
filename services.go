package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello world",
	})
}

func GetAllTasks(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, taskList)
}

func CreateTask(ctx *gin.Context) {
	var newTask Task

	if error := ctx.BindJSON(&newTask); error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": error.Error(),
		})
		return
	}

	newTask.Id = len(taskList) + 1
	taskList = append(taskList, newTask)

	ctx.JSON(http.StatusOK, newTask)
}

func GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, task := range taskList {
		if fmt.Sprintf("%d", task.Id) == id {
			ctx.JSON(http.StatusOK, task)
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task não encontrada"})
}

func DeleteTaskById(ctx *gin.Context) {
	id := ctx.Param("id")

	for index, task := range taskList {
		if fmt.Sprintf("%d", task.Id) == id {
			taskList = append(taskList[:index], taskList[index+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Task deletada com sucesso"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task não encontrada"})
}

func UpdateTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask Task

	if error := ctx.BindJSON(&updatedTask); error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": error.Error(),
		})
		return
	}

	for index, task := range taskList {
		if fmt.Sprintf("%d", task.Id) == id {
			updatedTask.Id = task.Id
			taskList[index] = updatedTask
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Task atualizada com sucesso",
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task não encontrada"})
}
