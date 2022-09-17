package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/samply/golang-fhir-models/fhir-models/fhir"
)

func getFHIRQuestionnaire(context *gin.Context) {
	content, err := ioutil.ReadFile("./example.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload fhir.Questionnaire
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
	// router.GET("/questionnaire", getQuestionnaire)
	router.GET("/questionnaire", getFHIRQuestionnaire)
	router.Run(":9090")
}
