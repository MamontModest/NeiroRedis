package redisService

import (
	"NeiroRedis/internal/models"
	"NeiroRedis/pkg/logger"
	"errors"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"net/http"
)

type resource struct {
	logger.Logger
	IRedisService
}

func InitHandlers(r *routing.RouteGroup, service IRedisService, log logger.Logger) {
	res := resource{IRedisService: service, Logger: log}

	r.Get("/", res.get)
	r.Post("/", res.set)
	r.Delete("/", res.delete)
}

func (res resource) get(ctx *routing.Context) error {
	return nil
}

func (res resource) set(ctx *routing.Context) error {
	message := new(models.Message)
	err := ctx.Read(&message)
	if err != nil {
		return errors.New("while reading form")
	}
	res.IRedisService.Set(*message)
	ctx.Response.WriteHeader(http.StatusNoContent)
	return nil
}

func (res resource) delete(ctx *routing.Context) error {
	return nil
}
