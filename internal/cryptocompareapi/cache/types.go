package cache

import "time"

const (
	CURRENCY_DATA = "CURRENCY_DATA"
)

type Storage interface {
	Set(key string, value interface{}, duration time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string) error
}
