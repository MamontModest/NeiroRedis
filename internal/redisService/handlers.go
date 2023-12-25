package redisService

import (
	"NeiroRedis/internal/models"
	"NeiroRedis/pkg/logger"
	"fmt"
	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type resource struct {
	logger.Logger
	IRedisService
}

func InitHandlers(r *routing.RouteGroup, service IRedisService, log logger.Logger) {
	res := resource{IRedisService: service, Logger: log}

	//if key is not exists get an error
	r.Get("/<key>", res.get)
	//if key is exist, update key
	r.Post("/", res.set)
	//if key is not exists get an error
	r.Delete("/<key>", res.delete)
}

func (res resource) get(ctx *routing.Context) error {
	//get param key
	key := ctx.Param("key")
	res.Logger.Info(fmt.Sprintf("%s, %v", key, ctx.Request))

	if key == "" {
		res.Logger.Error(fmt.Sprintf("cant validate key %v", ctx.Request))
		return ctx.WriteWithStatus(BadRequestResponse{Err: "cant validate key"}, http.StatusBadRequest)
	}
	object, err := res.IRedisService.Get(key)
	if err != nil {
		res.Logger.Error(err.Error())
		return ctx.WriteWithStatus(BadRequestResponse{Err: err.Error()}, http.StatusBadRequest)
	}

	return ctx.WriteWithStatus(object, http.StatusOK)
}

func (res resource) set(ctx *routing.Context) error {
	message := new(models.MessageSet)
	err := ctx.Read(&message)

	if err != nil {
		res.Logger.Error(fmt.Sprintf("cant parse form %v", ctx.Request))
		return ctx.WriteWithStatus(BadRequestResponse{Err: "cant parse form"}, http.StatusBadRequest)
	}
	res.Logger.Info(fmt.Sprintf("%v, %v", message, ctx.Request))

	validate := validator.New()
	err = validate.Struct(*message)
	if err != nil {
		res.Logger.Error(fmt.Sprintf("cant validate form %v", ctx.Request))
		return ctx.WriteWithStatus(BadRequestResponse{Err: "cant validate form"}, http.StatusBadRequest)
	}

	res.IRedisService.Set(*message)
	return ctx.WriteWithStatus(message, http.StatusCreated)
}

func (res resource) delete(ctx *routing.Context) error {
	//get param key
	key := ctx.Param("key")
	res.Logger.Info(fmt.Sprintf("%s, %v", key, ctx.Request))

	if key == "" {
		res.Logger.Error(fmt.Sprintf("cant validate key %v", ctx.Request))
		return ctx.WriteWithStatus(BadRequestResponse{Err: "cant validate key"}, http.StatusBadRequest)
	}
	err := res.IRedisService.Delete(key)
	if err != nil {
		res.Logger.Error(err.Error())
		return ctx.WriteWithStatus(BadRequestResponse{Err: err.Error()}, http.StatusBadRequest)
	}
	ctx.Response.WriteHeader(http.StatusNoContent)
	return nil
}

type BadRequestResponse struct {
	Err string `json:"err"`
}
