package ports

import "time"

type PortCache interface {
	AddKeyValue(key string, data interface{}, ex time.Duration) error
	GetWithKey(key string, st interface{}) error
	GetWithPrefix(prefix string, st interface{}) error
	DeleteByKey(key string) error
	DeleteByPrefix(prefix string) error
}
