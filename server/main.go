package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Element struct {
	Name          string   `json:"name"`
	Title         string   `json:"title"`
	Type          string   `json:"type"`
	Html          string   `json:"html"`
	IsRequired    bool     `json:"isRequired"`
	Choices       []string `json:"choices"`
	CorrectAnswer string   `json:"correctAnswer"`
	ChoicesOrder  string   `json:"choicesOrder"`
}

type Elements struct {
	Elements []Element `json:"elements"`
}

type SurveyJson struct {
	Title string     `json:"title"`
	Pages []Elements `json:"pages"`
}

func checkResourceType(r Questionnaire, rt string) {
	if r.ResourceType != rt {
		log.Fatal("Only FHIR Questionnaire resources are supported")

	}
	log.Println("ResourceType is Questionnaire")
}

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

		fmt.Println(q)
		var e Element
		e.Title = q.Text
		e.Name = q.LinkId

		if q.Type == "string" {
			e.Type = "text"
		}

		if q.Type == "choice" {
			e.Type = "radiogroup"
		}

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
