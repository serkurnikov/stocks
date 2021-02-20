package dal

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"stocks/internal/app"
	"stocks/internal/cryptocompareapi"
)

const (
	filePath = "/internal/dal/resourse_files/settings.yaml"
)

type ResourseData struct {}

func Init() app.ResourseData {
	return &ResourseData{}
}

func (r ResourseData) GetDataFromYamlResource() (params *cryptocompareapi.CurrencyParams) {
	currencyParams := cryptocompareapi.CurrencyParams{}

	err := yaml.Unmarshal(readFile(filePath), &currencyParams)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	return &currencyParams
}

func readFile(path string) []byte {
	pwd, error := os.Getwd()
	if error != nil {
		yamlFile, err := ioutil.ReadFile(pwd + path)
		if err != nil {
			return yamlFile
		}
	}
	return nil
}
