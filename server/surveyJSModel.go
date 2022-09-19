package main

type ChoicesByUrl struct {
	Url       string `json:"url"`
	ValueName string `json:"valueName"`
}
type Element struct {
	Name          string       `json:"name"`
	Title         string       `json:"title"`
	Type          string       `json:"type"`
	Html          string       `json:"html"`
	IsRequired    bool         `json:"isRequired"`
	Choices       []string     `json:"choices"`
	CorrectAnswer string       `json:"correctAnswer"`
	ChoicesOrder  string       `json:"choicesOrder"`
	ChoicesByUrl  ChoicesByUrl `json:"choicesByUrl"`
}

type Elements struct {
	Elements []Element `json:"elements"`
}

type GenderOption struct {
	Code   string `json:"code"`
	Diplay string `json:"display"`
}

type SurveyJson struct {
	Title string     `json:"title"`
	Pages []Elements `json:"pages"`
}
