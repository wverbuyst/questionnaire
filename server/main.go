package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getFHIRQuestionnaire(context *gin.Context) {
	content, err := ioutil.ReadFile("./example.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload Questionnaire
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	checkResourceType(payload, "Questionnaire")

	var surveyJson SurveyJson
	surveyJson.Title = payload.Title

	for _, q := range payload.Item {
		var e Element
		e.Title = q.Text
		e.Name = q.LinkId

		t, err := getType(q.Type)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		e.Type = t

		for _, o := range q.AnswerOption {
			e.Choices = append(e.Choices, o.ValueCoding.Code)
		}

		var es Elements
		es.Elements = append(es.Elements, e)
		surveyJson.Pages = append(surveyJson.Pages, es)
	}

	context.IndentedJSON(http.StatusOK, surveyJson)
}

func welcomeWorld(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"message": "hello world"})
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/", welcomeWorld)
	router.GET("/questionnaire", getFHIRQuestionnaire)
	router.Run(":9090")
}
