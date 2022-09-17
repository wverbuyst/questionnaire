package main

type Questionnaire struct {
	Id             string              `json:"id,omitempty"`
	ResourceType   string              `json:"resourceType"`
	ImplicitRules  string              `json:"implicitRules,omitempty"`
	Language       string              `json:"language,omitempty"`
	Url            string              `json:"url,omitempty"`
	Version        string              `json:"version,omitempty"`
	Name           string              `json:"name,omitempty"`
	Title          string              `json:"title,omitempty"`
	Date           string              `json:"date,omitempty"`
	Publisher      string              `json:"publisher,omitempty"`
	Description    string              `json:"description,omitempty"`
	Purpose        string              `json:"purpose,omitempty"`
	Copyright      string              `json:"copyright,omitempty"`
	ApprovalDate   string              `json:"approvalDate,omitempty"`
	LastReviewDate string              `json:"lastReviewDate,omitempty"`
	Item           []QuestionnaireItem `json:"item,omitempty"`
}

type QuestionnaireItem struct {
	Id             string         `json:"id,omitempty"`
	LinkId         string         `json:"linkId"`
	Definition     string         `json:"definition,omitempty"`
	Prefix         string         `json:"prefix,omitempty"`
	Text           string         `json:"text,omitempty"`
	Type           string         `json:"type"`
	Required       bool           `json:"required,omitempty"`
	Repeats        bool           `json:"repeats,omitempty"`
	ReadOnly       bool           `json:"readOnly,omitempty"`
	MaxLength      int            `json:"maxLength,omitempty"`
	AnswerValueSet string         `json:"answerValueSet,omitempty"`
	AnswerOption   []AnswerOption `json:"answerOption,omitempty"`
}

type AnswerOption struct {
	ValueCoding Code `bson:"valueCoding,omitempty" json:"valueCoding,omitempty"`
}

type Code struct {
	Code string `bson:"code,omitempty" json:"code,omitempty"`
}
