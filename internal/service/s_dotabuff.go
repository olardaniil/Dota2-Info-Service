package service

import (
	"dota2_info_service/internal/entity"
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

type DotabuffService struct {
}

func NewDotabuffService() *DotabuffService {
	return &DotabuffService{}
}

func (s *DotabuffService) GetHeroByName(heroName string) (entity.Hero, error) {
	var hero entity.Hero
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Mobile Safari/537.36"

	var lines []entity.Line
	// Получаем информацию о популярных линиях
	selector := "body > div.container-outer.seemsgood > div.skin-container > div.container-inner.container-inner-content > div.content-inner > div.row-12.with-sidebar > div.col-8 > section:nth-child(2) > article > table > tbody"
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			var line entity.Line
			line.Name = el.ChildText("td:nth-child(1)")
			line.Presence = el.ChildText("td:nth-child(2)")
			line.WinRate = el.ChildText("td:nth-child(3)")
			lines = append(lines, line)
		})
	})

	var strongOpponents []entity.StrongOpponent
	// Получаем информацию о 5 самых сильных противниках для героя
	selector = "body > div.container-outer.seemsgood > div.skin-container > div.container-inner.container-inner-content > div.content-inner > div.row-12.with-sidebar > div.col-8 > section:nth-child(11) > article > table > tbody"
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			if len(strongOpponents) == 5 {
				return
			}
			var strongOpponent entity.StrongOpponent
			strongOpponent.Name = el.ChildText("td:nth-child(2) > a")
			strongOpponent.Url = fmt.Sprintf("https://ru.dotabuff.com%s", el.ChildAttr("td:nth-child(2) > a", "href"))
			strongOpponent.WinRateOriginalHero = el.ChildText("td:nth-child(4)")
			strongOpponents = append(strongOpponents, strongOpponent)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	err := c.Visit("https://ru.dotabuff.com/heroes/" + heroName)
	if err != nil {
		return hero, err
	}

	hero = entity.Hero{
		Name:            heroName,
		PopularLines:    lines,
		StrongOpponents: strongOpponents,
	}

	return hero, nil
}
