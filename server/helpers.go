package main

import (
	"errors"
	"log"
)

func checkResourceType(r Questionnaire, rt string) {
	if r.ResourceType != rt {
		log.Fatal("Only FHIR Questionnaire resources are supported")
	}
}

func getTextForHtml(r Questionnaire) Elements {
	var es Elements
	var e Element
	if r.Text.Div != "" {
		e.Type = "html"
		e.Html = r.Text.Div
	}
	es.Elements = append(es.Elements, e)
	return es
}

func getType(t string) (string, error) {
	switch t {
	case "boolean":
		return "boolean", nil
	case "choice":
		return "radiogroup", nil
	case "date":
		return "text", nil
	case "dateTime":
		return "text", nil
	case "decimal":
		return "text", nil
	case "display":
		return "html", nil
	case "group":
		return "panel", nil
	case "integer":
		return "text", nil
	case "string":
		return "text", nil
	case "text":
		return "comment", nil
	case "time":
		return "text", nil
	case "url":
		return "text", nil
	default:
		return "", errors.New("Unhandled value[x] in answerOption")
	}
}
