package cryptocompareapi

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const (
	filePath = "/internal/cryptocompareapi/resourse_files/settings.yaml"
)

func GetDataFromYamlResource() (params *CurrencyParams) {
	currencyParams := CurrencyParams{}

	err := yaml.Unmarshal(readFile(filePath), &currencyParams)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	return &currencyParams
}

func readFile(path string) []byte {
	pwd, error := os.Getwd()
	if error == nil {
		yamlFile, err := ioutil.ReadFile(pwd + path)
		if err == nil {
			return yamlFile
		}
	}
	return nil
}
