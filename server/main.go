package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Element struct {
	Name  string
	Title string
	Type  string
}

type Questionnaire struct {
	Elements []Element
}

func getQuestionnaire(context *gin.Context) {
	content, err := ioutil.ReadFile("./questionnaire.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload Questionnaire
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	context.IndentedJSON(http.StatusOK, payload)
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
