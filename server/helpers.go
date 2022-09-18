package main

import "log"

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
