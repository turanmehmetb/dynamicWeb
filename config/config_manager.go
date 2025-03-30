package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"dynamicWeb/models"

	"gopkg.in/yaml.v3"
)

const dataPath = "config_files"

func getFilePath(id string) string {
	return filepath.Join(dataPath, id+".yaml")
}

func SaveConfiguration(config *models.Configuration) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(getFilePath(config.ID), data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetConfiguration(id string) (*models.Configuration, error) {
	filePath := getFilePath(id)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config models.Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func GetAllConfigurations() ([]models.Configuration, error) {
	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		return nil, err
	}

	var configurations []models.Configuration
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".yaml" {
			id := file.Name()[:len(file.Name())-5]
			config, err := GetConfiguration(id)
			if err == nil {
				configurations = append(configurations, *config)
			}
		}
	}

	return configurations, nil
}

func UpdateConfiguration(id string, updatedConfig *models.Configuration) error {
	_, err := GetConfiguration(id)
	if err != nil {
		return err
	}

	updatedConfig.ID = id
	return SaveConfiguration(updatedConfig)
}

func DeleteConfiguration(id string) error {
	filePath := getFilePath(id)
	return os.Remove(filePath)
}
