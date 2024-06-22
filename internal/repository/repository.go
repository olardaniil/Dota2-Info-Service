package repository

import (
	"dota2_info_service/internal/entity"
	"dota2_info_service/internal/repository/redis"
	"dota2_info_service/pkg/database"
)

type CacheHero interface {
	GetHeroByName(heroName string) (entity.Hero, error)
	SetHero(hero entity.Hero) error
}

type Repository struct {
	CacheHero
}

func NewRepository(dbRedis *database.Redis) *Repository {
	return &Repository{
		CacheHero: redis.NewCacheHeroRepo(dbRedis),
	}
}
