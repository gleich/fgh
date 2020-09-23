package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Write to a yaml file
func WriteYAML(data interface{}, fName string) error {
	YAMLContent, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fName, YAMLContent, 0644)
	return err
}

// Read from a yaml file
func ReadYAML(fName string, data interface{}) error {
	byteContent, err := ioutil.ReadFile(fName)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(byteContent, data)
	return err
}
