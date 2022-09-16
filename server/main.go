package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getQuestionnaire(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "here's a questionnaire for ya"})
}

func welcomeWorld(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "hello world"})
}

func main() {

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", welcomeWorld)
	router.GET("/questionnaire", getQuestionnaire)
	router.Run(":9090")
}
