package main

import (
	"NeiroRedis/internal/cache"
	"NeiroRedis/internal/heap"
	"NeiroRedis/internal/redisService"
	"NeiroRedis/pkg/logger"
	"fmt"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	"github.com/go-ozzo/ozzo-routing/v2/cors"
	"net/http"
)

const (
	logPath                    = "./logger/log.txt"
	MaxObjects                 = 15
	LimitDeletedObjectsForOnce = 3
	TimeExp                    = 60
)

func main() {
	Logger, err := logger.InitLogger(logPath)
	defer Logger.Close()
	if err != nil {
		panic(err)
	}

	address := fmt.Sprintf(":%v", "8000")
	router := routing.New()
	router.Use(
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.AllowAll),
	)
	redisService.InitHandlers(
		router.Group(""),
		redisService.RedisService{
			Heap: &heap.Heap{
				Objects:     make([]*heap.Object, 0),
				IndexObject: make(map[string]int),
			},
			Cache:                      cache.Cache{Storage: make(map[string]interface{})},
			TimeExp:                    TimeExp,
			MaxObjects:                 MaxObjects,
			LimitDeletedObjectsForOnce: LimitDeletedObjectsForOnce,
		},
		Logger,
	)
	hs := &http.Server{
		Addr:    address,
		Handler: router,
	}
	err = hs.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
