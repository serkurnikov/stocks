package cryptocompareapi

import (
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"stocks/internal/cryptocompareapi/cache"
	"stocks/internal/cryptocompareapi/cache/memory"
	"stocks/internal/dao"
	"strings"
	"time"
)

const (
	cachedTime  = 5 * time.Minute
	clearedTime = 1 * time.Second
	timeUpdate  = 2 * time.Second
	baseURL     = "https://min-api.cryptocompare.com/data/pricemultifull?"
)

type Api interface {
	GetCurrencyPrice(params *CurrencyParams) (*gabs.Container, error)
	UpdateCurrency()
}

type api struct {
	client         *http.Client
	storage        cache.Storage
	currencyParams *CurrencyParams
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

func (c *api) GetCurrencyPrice(params *CurrencyParams) (*gabs.Container, error) {
	//Get Data from Cash
	cacheData, exist := c.storage.Get(cache.CURRENCY_DATA)
	jsonData, _ := json.Marshal(cacheData)

	if exist {
		return gabs.ParseJSON(jsonData)
	}

	fsyms := strings.Split(params.Fsyms, splitS)
	tsyms := strings.Split(params.Tsyms, splitS)

	p, _ := structToMap(params)
	body, _ := c.req(p)
	jsonParsed, _ := gabs.ParseJSON(body)
	outPut := gabs.New()

	for i := 0; i < len(fsyms); i++ {
		for j := 0; j < len(tsyms); j++ {
			var basePath = splitPoint + fsyms[i] + splitPoint + tsyms[j]
			var rawPath = RAW + basePath
			var displayPath = DISPLAY + basePath

			for key, child := range jsonParsed.Path(rawPath).ChildrenMap() {
				if isExist(key) {
					outPut.SetP(fmt.Sprintf("%v", child.Data().(interface{})), rawPath+splitPoint+key)
				}
			}

			for key, child := range jsonParsed.Path(displayPath).ChildrenMap() {
				if isExist(key) {
					outPut.SetP(fmt.Sprintf("%v", child.Data().(interface{})), displayPath+splitPoint+key)
				}
			}

		}
	}
	println("Update Data", outPut.String())
	return outPut, nil
}

func isExist(key string) bool {
	if key == dao.CurrencyValuesConstants[key] {
		return true
	} else {
		return false
	}
}

func NewCryptoCompare(params *CurrencyParams) Api {
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
		currencyParams: params,
	}
}

func (c *api) UpdateCurrency() {
	ticker := time.NewTicker(timeUpdate)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			//Save Data to Cash
			result, _ := c.GetCurrencyPrice(c.currencyParams)
			c.storage.Set(cache.CURRENCY_DATA, result, clearedTime)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
