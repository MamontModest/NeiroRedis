package main

import (
	"NeiroRedis/internal/redisService"
	"NeiroRedis/pkg/logger"
	"fmt"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	"github.com/go-ozzo/ozzo-routing/v2/cors"
	"net/http"
)

const logPath = "./logger/log.txt"

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
	redisService.InitHandlers(router.Group(""), redisService.RedisService{}, Logger)
	hs := &http.Server{
		Addr:    address,
		Handler: router,
	}
	err = hs.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
