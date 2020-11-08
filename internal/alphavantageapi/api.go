package alphavantageapi

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"stocks/internal/alphavantageapi/cache"
	"stocks/internal/alphavantageapi/cache/memory"
	"strings"
	"time"
)

const (
	cachedTime  = 5 * time.Minute
	clearedTime = 30 * time.Minute
	baseURL     = "https://www.alphavantage.co/query?"
	apiKey      = "KI4JR6Z6SBQSZTNA"
)

type Api interface {
	GetTimeSeriesIntraday(params TimeSeriesIntradayParams) (*TimeSeriesIntraday, error)
}

type api struct {
	client  *http.Client
	storage cache.Storage
}

func getEncodedURL(params map[string]string) string {
	base, err := url.Parse(baseURL)
	if err != nil {
		return baseURL
	}

	q := url.Values{}
	for k, v := range params {
		if len(v) != 0 {
			q.Add(strings.ToLower(k), v)
		}
	}
	q.Add("apikey", apiKey)

	base.RawQuery = q.Encode()
	return base.String()
}

func (c *api) req(params map[string]string) ([]byte, error) {

	req, err := c.client.Get(getEncodedURL(params))

	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *api) GetTimeSeriesIntraday(params TimeSeriesIntradayParams) (*TimeSeriesIntraday, error) {
	/*cachedRates, ok := c.storage.Get(timeSeriesIntraday)
	if ok {
		//return object from cash
	}*/

	var tsi TimeSeriesIntraday
	p, _ := structToMap(params)
	body, _ := c.req(p)
	if err := json.Unmarshal(body, &tsi); err != nil {
		//return object from database
		return nil, err
	}

	c.storage.Set(timeSeriesIntraday, tsi, 0)
	return &tsi, nil
}

func NewAlphaVantage() Api {
	defaultTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:        20,
		MaxIdleConnsPerHost: 20,
		TLSHandshakeTimeout: 15 * time.Second,
	}

	client := &http.Client{
		Transport: defaultTransport,
		Timeout:   15 * time.Second,
	}

	return &api{
		client:  client,
		storage: memory.InitCash(cachedTime, clearedTime),
	}
}
