package alphavantageapi

import (
	"encoding/json"
)

//API Parameters TIME SERIES INTRADAY
/*
@required: function. For example: function=TIME_SERIES_INTRADAY
@required: symbol. For example: symbol=IBM
@required: interval. The following values are supported: 1min, 5min, 15min, 30min, 60min

@optional: adjusted. By default, adjusted=true and the output time series is adjusted by
historical split and dividend events. Set adjusted=false to query raw (as-traded) intraday values.

@optional: outputsize. Strings compact and full are accepted with the following specifications:
compact returns only the latest 100 data points in the intraday time series; full returns the
full-length intraday time series. The "compact" option is recommended if you would like to
reduce the data size of each API call.

@optional: datatype. For example: datatype=json. Strings json and csv are accepted with the
following specifications: json returns the intraday time series in JSON format; csv returns
the time series as a CSV (comma separated value) file.
*/

const (
	timeSeriesIntraday = "TIME_SERIES_INTRADAY"
)

type (
	TimeSeriesIntradayParams struct {
		Function, Symbol, Interval, Adjusted, Outputsize, Datatype string
	}

	TimeSeriesIntraday struct {
		MetaData struct {
			Information   string `json:"1. Information"`
			Symbol        string `json:"2. Symbol"`
			LastRefreshed string `json:"3. Last Refreshed"`
			Interval      string `json:"4. Interval"`
			OutputSize    string `json:"5. Output Size"`
			TimeZone      string `json:"6. Time Zone"`
		} `json:"Meta Data"`
		TimeSeries5Min struct{} `json:"Time Series (5min)"`
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
