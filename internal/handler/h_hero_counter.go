package handler

import (
	"dota2_info_service/internal/entity"
	"net/http"
)

func (h *Handler) HeroCounterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.GetHeroInfoCounter(w, r)
	}
}

// GetHeroInfoCounter
// @Summary Получить информацию о персонажах, против которых заданные герой играет плохо
// @Tags heroes
// @Accept json
// @Produce json
// @Param hero path string true "Имя героя"
// @Success 200 {object} entity.HeroInfoCounterResponse
// @Failure 400 {object} Error
// @Failure 404 {object} Error
// @Router /counter/{hero} [get]
func (h *Handler) GetHeroInfoCounter(w http.ResponseWriter, r *http.Request) {
	// Получаем имя героя
	var hero = new(entity.Hero)
	hero.Name = r.URL.Path[len("/counter/"):]
	// Валидация
	err := hero.Validate()
	if err != nil {
		h.response.SendError(w, 400, err.Error())
		return
	}
	// Получаем информацию о герое
	heroInfo, err := h.service.Hero.GetHeroByName(hero.Name)
	if err != nil {
		if err.Error() == "Not Found" {
			h.response.SendError(w, 404, "Герой не найден")
			return
		}
		h.response.SendError(w, 400, err.Error())
		return
	}
	hero.StrongOpponents = heroInfo.StrongOpponents
	// Отправляем ответ
	h.response.SendJSON(w, 200, hero)
}
