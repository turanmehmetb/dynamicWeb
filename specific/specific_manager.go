package specific

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"dynamicWeb/models"

	"gopkg.in/yaml.v3"
)

const dataPath = "specific_files"

func getFilePath(id string) string {
	return filepath.Join(dataPath, id+".yaml")
}

func SaveSpecific(config *models.Specific) error {
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

func GetSpecific(id string) (*models.Specific, error) {
	filePath := getFilePath(id)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config models.Specific
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return GetOnlyIds(&config), nil
}

func GetAllSpecifics() ([]models.Specific, error) {
	files, err := ioutil.ReadDir(dataPath)
	if err != nil {
		return nil, err
	}

	var configurations []models.Specific
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".yaml" {
			id := file.Name()[:len(file.Name())-5]
			config, err := GetSpecific(id)
			if err == nil {
				configurations = append(configurations, *config)
			}
		}
	}

	return configurations, nil
}

func UpdateSpecific(id string, updatedConfig *models.Specific) error {
	_, err := GetSpecific(id)
	if err != nil {
		return err
	}

	updatedConfig.ID = id
	return SaveSpecific(updatedConfig)
}

func DeleteSpecific(id string) error {
	filePath := getFilePath(id)
	return os.Remove(filePath)
}

func GetOnlyIds(config *models.Specific) *models.Specific {
	for key, values := range config.Datasource.Hosts {
		for i, val := range values {
			config.Datasource.Hosts[key][i] = strings.TrimSuffix(val, ".yaml")
		}
	}
	for key, values := range config.Datasource.Urls {
		for i, val := range values {
			config.Datasource.Urls[key][i] = strings.TrimSuffix(val, ".yaml")
		}
	}
	for key, values := range config.Datasource.Pages {
		for i, val := range values {
			config.Datasource.Pages[key][i] = strings.TrimSuffix(val, ".yaml")
		}
	}

	return config
}
