package redisService

import (
	"NeiroRedis/internal/cache"
	"NeiroRedis/internal/heap"
	"NeiroRedis/internal/models"
	"errors"
)

type RedisService struct {
	Heap  heap.IHeap
	Cache cache.ICache
}

type IRedisService interface {
	Get(message models.Message) (interface{}, error)
	Set(message models.Message)
	Delete(message models.Message) error
}

func (r RedisService) Get(message models.Message) (interface{}, error) {
	if value, ok := r.Cache.Get(message.Key); ok {
		return value, nil
	}
	return nil, errors.New("not data")
}

func (r RedisService) Set(message models.Message) {

}

func (r RedisService) Delete(message models.Message) error {
	return nil
}
