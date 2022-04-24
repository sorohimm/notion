package services

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"math"
	"notion/internal/models"
)

type HackersService struct {
	Log         *zap.SugaredLogger
	RedisClient *redis.Client
}

func (s *HackersService) GetHackers() []models.Hackers {
	result, err := s.RedisClient.ZRangeWithScores(context.Background(), "hackers", 0, math.MaxInt64).Result() // возвращаем всё тк интервал не задается
	if err != nil {
		s.Log.Warn(err.Error())
	}

	var hackersList []models.Hackers
	for _, zItem := range result {
		hackersList = append(hackersList, models.Hackers{Score: zItem.Score, Name: fmt.Sprintf("%v", zItem.Member)})
	}

	return hackersList
}
