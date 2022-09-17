package main

import "log"

func checkResourceType(r Questionnaire, rt string) {
	if r.ResourceType != rt {
		log.Fatal("Only FHIR Questionnaire resources are supported")

	}
	log.Println("ResourceType is Questionnaire")
}
