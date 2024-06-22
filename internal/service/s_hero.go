package service

import (
	"dota2_info_service/internal/entity"
	"dota2_info_service/internal/repository"
)

type HeroService struct {
	cacheHeroRepo repository.CacheHero
	dotabuff      Dotabuff
}

func NewHeroService(cacheHeroRepo repository.CacheHero, dotabuff Dotabuff) *HeroService {
	return &HeroService{
		cacheHeroRepo: cacheHeroRepo,
		dotabuff:      dotabuff,
	}
}

func (s *HeroService) GetHeroByName(heroName string) (entity.Hero, error) {
	// Смотрим в кэше
	hero, err := s.cacheHeroRepo.GetHeroByName(heroName)
	if err != nil {
		if err.Error() != "redis: nil" {
			return hero, err
		}
	}
	// Если есть в кэше, то возвращаем
	if hero.Name != "" {
		return hero, nil
	}
	// Если нет в кэше, то получаем из dotabuff service
	hero, err = s.dotabuff.GetHeroByName(heroName)
	if err != nil {
		return hero, err
	}
	// Записываем в кэш
	err = s.cacheHeroRepo.SetHero(hero)
	if err != nil {
		return hero, err
	}

	return hero, nil
}
