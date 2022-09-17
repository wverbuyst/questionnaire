package main

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
