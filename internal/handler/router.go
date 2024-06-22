package handler

import (
	_ "dota2_info_service/docs"
	"dota2_info_service/internal/service"
	"fmt"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

type Handler struct {
	service  *service.Service
	response *ResponseFactory
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service:  service,
		response: NewResponseFactory(),
	}
}

func (h *Handler) Run(port string) error {
	// api
	http.HandleFunc("/info/", h.HeroInfoHandler)
	http.HandleFunc("/counter/", h.HeroCounterHandler)
	// swagger
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	//
	fmt.Println("Сервис запущен")
	fmt.Println("http://localhost:" + port + "/swagger/")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		return err
	}
	return nil
}
