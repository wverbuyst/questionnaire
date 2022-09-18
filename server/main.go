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

	html := getTextForHtml(payload)
	surveyJson.Pages = append(surveyJson.Pages, html)

	for _, i := range payload.Item {
		var e Element
		var es Elements

		e.Title = i.Text
		e.Name = i.LinkId

		t, err := getType(i.Type)
		if err != nil {
			log.Fatal("Error: ", err)
		}
		e.Type = t

		for _, o := range i.AnswerOption {
			e.Choices = append(e.Choices, o.ValueCoding.Code)
		}

		es.Elements = append(es.Elements, e)

		processNestedItems(i, &es)

		log.Println("M", es)
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
