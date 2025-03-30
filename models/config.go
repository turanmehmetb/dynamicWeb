package models

import "github.com/google/uuid"

type Action struct {
	Type       string `yaml:"type" json:"type"`
	Selector   string `yaml:"selector,omitempty" json:"selector,omitempty"`
	Element    string `yaml:"element,omitempty" json:"element,omitempty"`
	NewElement string `yaml:"newElement,omitempty" json:"newElement,omitempty"`
	Position   string `yaml:"position,omitempty" json:"position,omitempty"`
	Target     string `yaml:"target,omitempty" json:"target,omitempty"`
	OldValue   string `yaml:"oldValue,omitempty" json:"oldValue,omitempty"`
	NewValue   string `yaml:"newValue,omitempty" json:"newValue,omitempty"`
}

type Configuration struct {
	ID      string   `yaml:"id" json:"id"`
	Actions []Action `yaml:"actions" json:"actions"`
}

func NewConfiguration(actions []Action) *Configuration {
	return &Configuration{
		ID:      uuid.New().String(),
		Actions: actions,
	}
}
