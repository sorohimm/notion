package infrastructure

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"notion/internal/controllers"
	"notion/internal/services"
)

type IInjector interface {
	InjectBalanceController() controllers.HackersController
}

var env *environment

type environment struct {
	logger      *zap.SugaredLogger
	redisClient *redis.Client
}

func (e *environment) InjectBalanceController() controllers.HackersController {
	return controllers.HackersController{
		Log: e.logger,
		RedisService: &services.HackersService{
			Log:         e.logger,
			RedisClient: e.redisClient,
		},
		Validator: validator.New(),
	}
}

func Injector(log *zap.SugaredLogger, redisClient *redis.Client) (IInjector, error) {
	env = &environment{
		logger:      log,
		redisClient: redisClient,
	}

	return env, nil
}
