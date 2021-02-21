package cryptocompareapi

import (
	"encoding/json"
)

const (
	RAW     = "RAW"
	DISPLAY = "DISPLAY"
	splitS     = ","
	splitPoint = "."
)

type (
	CurrencyParams struct {
		Fsyms string `yaml:"fsyms"`
		Tsyms string `yaml:"tsyms"`
	}
)

/*Convert struct to map*/
func structToMap(data interface{}) (map[string]string, error) {

	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	mapData := make(map[string]string)
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}

	return mapData, nil
}
