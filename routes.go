package main

import "github.com/gin-gonic/gin"

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

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", MainRoute)
	router.GET("/tasks", GetAllTasks)
	router.POST("/tasks", CreateTask)
	router.GET("/tasks/:id", GetTaskById)
	router.DELETE("/tasks/:id", DeleteTaskById)
	router.PUT("/tasks/:id", UpdateTaskById)
}
