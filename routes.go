package main

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", MainRoute)
	router.GET("/tasks", GetAllTasks)
	router.POST("/tasks", CreateTask)
	router.GET("/tasks/:id", GetTaskById)
	router.DELETE("/tasks/:id", DeleteTaskById)
	router.PUT("/tasks/:id", UpdateTaskById)
}
