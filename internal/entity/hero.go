package entity

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Hero struct {
	Name            string           `json:"name,omitempty"`
	PopularLines    []Line           `json:"popular_lines,omitempty"`
	StrongOpponents []StrongOpponent `json:"strong_opponents,omitempty"`
}

func (h *Hero) Validate() error {
	// Проверка имени
	if h.Name == "" {
		return fmt.Errorf("имя героя не может быть пустым")
	}
	// Берём только первую часть до "/"
	h.Name = strings.Split(h.Name, "/")[0]

	return nil
}

func (h *Hero) MarshalBinary() ([]byte, error) {
	return json.Marshal(h)
}

func (h *Hero) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &h)
}

type HeroInfoResponse struct {
	Name         string `json:"name,omitempty"`
	PopularLines []Line `json:"popular_lines,omitempty"`
}

type HeroInfoCounterResponse struct {
	Name            string           `json:"name,omitempty"`
	StrongOpponents []StrongOpponent `json:"strong_opponents,omitempty"`
}
