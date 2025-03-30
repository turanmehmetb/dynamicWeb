package models

import "github.com/google/uuid"

type Datasource struct {
	Pages map[string][]string `yaml:"pages" json:"pages"`
	Urls  map[string][]string `yaml:"urls" json:"urls"`
	Hosts map[string][]string `yaml:"hosts" json:"hosts"`
}

type Specific struct {
	ID         string     `yaml:"id" json:"id"`
	Datasource Datasource `yaml:"datasource" json:"datasource"`
}

func NewSpecific(datasource Datasource) *Specific {
	return &Specific{
		ID:         uuid.New().String(),
		Datasource: datasource,
	}
}
