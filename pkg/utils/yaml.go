package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Write to a yaml file
func WriteYAML(data interface{}, fName string) CtxErr {
	YAMLContent, err := yaml.Marshal(data)
	if err != nil {
		return CtxErr{
			Error:   err,
			Context: "Failed to convert data to yaml",
		}
	}

	parentDir := filepath.Dir(fName)
	err = os.MkdirAll(parentDir, 0777)
	if err != nil {
		return CtxErr{
			Error:   err,
			Context: "Failed to make make directory: " + parentDir,
		}
	}

	err = ioutil.WriteFile(fName, YAMLContent, 0644)
	return CtxErr{
		Error:   err,
		Context: "Failed to write data to " + fName,
	}
}

// Read from a yaml file
func ReadYAML(fName string, data interface{}) CtxErr {
	byteContent, err := ioutil.ReadFile(fName)
	if err != nil {
		return CtxErr{
			Error:   err,
			Context: "Failed to read from file",
		}
	}
	err = yaml.Unmarshal(byteContent, data)
	return CtxErr{
		Error:   err,
		Context: "Failed to parse yaml file",
	}
}
