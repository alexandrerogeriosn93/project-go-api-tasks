package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

var taskList = []Task{
	{Id: 1, Title: "Estudar Go"},
	{Id: 2, Title: "Estudar Python"},
	{Id: 3, Title: "Praticar Git"},
	{Id: 4, Title: "Ler sobre Docker"},
	{Id: 5, Title: "Aprender SQL"},
	{Id: 6, Title: "Estudar Algoritmos"},
	{Id: 7, Title: "Revisar código de projeto"},
	{Id: 8, Title: "Explorar novas bibliotecas Go"},
	{Id: 9, Title: "Estudar estrutura de dados"},
	{Id: 10, Title: "Ler artigo sobre Inteligência Artificial"},
	{Id: 11, Title: "Desenvolver API com Go"},
	{Id: 12, Title: "Criar projeto de CRUD em Python"},
	{Id: 13, Title: "Explorar novas funcionalidades do Docker"},
	{Id: 14, Title: "Estudar Design Patterns em Go"},
	{Id: 15, Title: "Analisar código open-source no GitHub"},
	{Id: 16, Title: "Praticar testes automatizados com Go"},
	{Id: 17, Title: "Ler sobre arquitetura de software"},
	{Id: 18, Title: "Estudar sobre computação distribuída"},
	{Id: 19, Title: "Configurar ambiente de desenvolvimento em Docker"},
	{Id: 20, Title: "Criar projeto de microserviço com Go"},
	{Id: 21, Title: "Aprender sobre UX/UI Design"},
	{Id: 22, Title: "Explorar ferramentas de prototipagem como Figma"},
	{Id: 23, Title: "Ler sobre marketing digital e SEO"},
	{Id: 24, Title: "Estudar criação de conteúdo para mídias sociais"},
	{Id: 25, Title: "Aprender desenvolvimento para dispositivos móveis"},
	{Id: 26, Title: "Desenvolver uma aplicação web com React"},
	{Id: 27, Title: "Estudar JavaScript avançado"},
	{Id: 28, Title: "Ler sobre segurança em desenvolvimento web"},
	{Id: 29, Title: "Fazer curso de Machine Learning básico"},
	{Id: 30, Title: "Estudar sobre computação quântica"},
	{Id: 31, Title: "Aprender sobre Blockchain e criptomoedas"},
	{Id: 32, Title: "Explorar ferramentas de CI/CD como Jenkins"},
	{Id: 33, Title: "Estudar teoria de redes e protocolos de comunicação"},
	{Id: 34, Title: "Ler sobre desenvolvimento ágil e Scrum"},
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	router.GET("/tasks", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, taskList)
	})

	router.POST("/tasks", func(ctx *gin.Context) {
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
	})

	router.GET("/tasks/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		for _, task := range taskList {
			if fmt.Sprintf("%d", task.Id) == id {
				ctx.JSON(http.StatusOK, task)
				return
			}
		}

		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task não encontrada"})
	})

	router.DELETE("/tasks/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		for index, task := range taskList {
			if fmt.Sprintf("%d", task.Id) == id {
				taskList = append(taskList[:index], taskList[index+1:]...)
				ctx.JSON(http.StatusOK, gin.H{"message": "Task deletada com sucesso"})
				return
			}
		}

		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task não encontrada"})
	})

	router.PUT("/tasks/:id", func(ctx *gin.Context) {
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
	})

	router.Run(":3000")
}
