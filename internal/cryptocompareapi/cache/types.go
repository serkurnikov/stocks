package cache

import "time"

const (
	CURRENCY_PARAMS = "CURRENCY_PARAMS"
)

type Storage interface {
	Set(key string, value interface{}, duration time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string) error
}
