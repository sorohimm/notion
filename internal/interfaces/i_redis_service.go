package interfaces

import "notion/internal/models"

type IRedisService interface {
	GetHackers() []models.Hackers
}
