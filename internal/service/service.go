package service

import (
	"dota2_info_service/internal/entity"
	"dota2_info_service/internal/repository"
)

type Hero interface {
	GetHeroByName(heroName string) (entity.Hero, error)
}

type Dotabuff interface {
	GetHeroByName(heroName string) (entity.Hero, error)
}

type Service struct {
	Hero
	Dotabuff
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Hero:     NewHeroService(repo, NewDotabuffService()),
		Dotabuff: NewDotabuffService(),
	}
}
