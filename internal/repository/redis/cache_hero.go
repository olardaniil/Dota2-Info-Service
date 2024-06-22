package redis

import (
	"dota2_info_service/internal/entity"
	"dota2_info_service/pkg/database"
	"fmt"
	"time"
)

type CacheHeroRepo struct {
	redis *database.Redis
}

func NewCacheHeroRepo(redis *database.Redis) *CacheHeroRepo {
	return &CacheHeroRepo{
		redis: redis,
	}
}

func (c *CacheHeroRepo) GetHeroByName(heroName string) (entity.Hero, error) {
	key := fmt.Sprintf("hero:%s", heroName)
	hero := entity.Hero{}
	err := c.redis.Client.Get(key).Scan(&hero)
	if err != nil {
		return entity.Hero{}, err
	}

	return hero, nil
}

func (c *CacheHeroRepo) SetHero(hero entity.Hero) error {
	key := fmt.Sprintf("hero:%s", hero.Name)
	heroByte, err := hero.MarshalBinary()
	if err != nil {
		return err
	}
	err = c.redis.Client.Set(key, heroByte, time.Hour*24).Err()
	if err != nil {
		return err
	}
	return nil
}
