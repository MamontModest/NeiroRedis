package redisService

import (
	"NeiroRedis/internal/cache"
	"NeiroRedis/internal/heap"
	"NeiroRedis/internal/models"
	"errors"
	"fmt"
	"time"
)

type RedisService struct {
	//Куча для выбора приоритета на удаление
	Heap heap.IHeap
	//Кэш для хранения объектов
	Cache cache.ICache
	//Время жизни объекта
	TimeExp int64
	//Максимальное кол-во объектов, если 0, то кол-во объектов не ограниченно
	MaxObjects uint64
	//Максимальное кол-во объектов, которые можно удалить за раз
	LimitDeletedObjectsForOnce int
}

type IRedisService interface {
	Get(key string) (interface{}, error)
	Set(message models.MessageSet)
	Delete(key string) error
}

func (r RedisService) Get(key string) (interface{}, error) {
	r.deleteTimeLimitedObjects()
	if value, ok := r.Cache.Get(key); ok {
		updateObject := &heap.Object{
			Key: key,
			Exp: time.Now().Unix() + r.TimeExp,
		}
		r.Heap.ChangeObject(updateObject)
		return value, nil
	}
	return nil, errors.New(fmt.Sprintf("not data with key %s", key))
}

func (r RedisService) Set(message models.MessageSet) {
	r.deleteTimeLimitedObjects()
	if _, ok := r.Cache.Get(message.Key); ok {
		r.Cache.Update(message.Key, message.Value)
		updateObject := &heap.Object{
			Key: message.Key,
			Exp: time.Now().Unix() + r.TimeExp,
		}
		r.Heap.ChangeObject(updateObject)
		return
	}
	r.Cache.Set(message.Key, message.Value)
	newObject := &heap.Object{
		Key: message.Key,
		Exp: time.Now().Unix() + r.TimeExp,
	}
	r.Heap.Push(newObject)
	r.deleteMaxObjects()
	return

}

func (r RedisService) Delete(key string) error {
	r.deleteTimeLimitedObjects()
	if ok := r.Cache.Delete(key); ok {
		if r.Heap.DeleteFromMiddle(key) {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("not data with key %s", key))
}

func (r RedisService) deleteTimeLimitedObjects() {
	for count := 0; count < r.LimitDeletedObjectsForOnce && r.Cache.Len() != 0; count++ {
		object, _ := r.Heap.GetLastItem()
		if object.Exp < time.Now().Unix() {
			object, _ = r.Heap.Pop()
			r.Cache.Delete(object.Key)
		}
	}
}

func (r RedisService) deleteMaxObjects() {
	for r.Cache.Len() > int(r.MaxObjects) && r.MaxObjects != 0 {
		object, _ := r.Heap.Pop()
		r.Cache.Delete(object.Key)
	}
}
