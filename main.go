package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

var taskList = []Task{
	{Id: 305, Title: "Estudar Go"},
	{Id: 123, Title: "Estudar Python"},
	{Id: 758, Title: "Praticar Git"},
	{Id: 419, Title: "Ler sobre Docker"},
	{Id: 213, Title: "Aprender SQL"},
	{Id: 861, Title: "Estudar Algoritmos"},
	{Id: 132, Title: "Revisar código de projeto"},
	{Id: 556, Title: "Explorar novas bibliotecas Go"},
	{Id: 727, Title: "Estudar estrutura de dados"},
	{Id: 832, Title: "Ler artigo sobre Inteligência Artificial"},
	{Id: 409, Title: "Desenvolver API com Go"},
	{Id: 986, Title: "Criar projeto de CRUD em Python"},
	{Id: 530, Title: "Explorar novas funcionalidades do Docker"},
	{Id: 845, Title: "Estudar Design Patterns em Go"},
	{Id: 614, Title: "Analisar código open-source no GitHub"},
	{Id: 431, Title: "Praticar testes automatizados com Go"},
	{Id: 372, Title: "Ler sobre arquitetura de software"},
	{Id: 998, Title: "Estudar sobre computação distribuída"},
	{Id: 501, Title: "Configurar ambiente de desenvolvimento em Docker"},
	{Id: 291, Title: "Criar projeto de microserviço com Go"},
	{Id: 694, Title: "Aprender sobre UX/UI Design"},
	{Id: 520, Title: "Explorar ferramentas de prototipagem como Figma"},
	{Id: 307, Title: "Ler sobre marketing digital e SEO"},
	{Id: 712, Title: "Estudar criação de conteúdo para mídias sociais"},
	{Id: 840, Title: "Aprender desenvolvimento para dispositivos móveis"},
	{Id: 640, Title: "Desenvolver uma aplicação web com React"},
	{Id: 921, Title: "Estudar JavaScript avançado"},
	{Id: 498, Title: "Ler sobre segurança em desenvolvimento web"},
	{Id: 563, Title: "Fazer curso de Machine Learning básico"},
	{Id: 142, Title: "Estudar sobre computação quântica"},
	{Id: 811, Title: "Aprender sobre Blockchain e criptomoedas"},
	{Id: 365, Title: "Explorar ferramentas de CI/CD como Jenkins"},
	{Id: 276, Title: "Estudar teoria de redes e protocolos de comunicação"},
	{Id: 588, Title: "Ler sobre desenvolvimento ágil e Scrum"},
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

	router.Run(":3000")
}
