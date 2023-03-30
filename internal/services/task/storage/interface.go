package storage

type Storage interface {
	GetByKey(key string) (interface{}, bool)
	Set(key string, data interface{})
}
