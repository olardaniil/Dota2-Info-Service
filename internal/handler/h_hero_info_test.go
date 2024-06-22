package handler

import (
	"dota2_info_service/internal/repository"
	"dota2_info_service/internal/service"
	"dota2_info_service/pkg/database"
	"github.com/alicebob/miniredis/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_GetHeroInfo(t *testing.T) {
	type fields struct {
		service  *service.Service
		response *ResponseFactory
	}
	type args struct {
		urlPath string
	}
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	clientRedis, err := database.NewRedis(mr.Addr(), "", 0)
	if err != nil {
		panic(err)
	}
	repo := repository.NewRepository(clientRedis)
	serv := service.NewService(repo)

	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStatusCode int
	}{
		{
			name: "Positive",
			fields: fields{
				service:  serv,
				response: NewResponseFactory(),
			},
			args: args{
				urlPath: "bristleback",
			},
			wantStatusCode: http.StatusOK,
		},
		{
			name: "Positive-2",
			fields: fields{
				service:  serv,
				response: NewResponseFactory(),
			},
			args: args{
				urlPath: "brist",
			},
			wantStatusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				service:  tt.fields.service,
				response: tt.fields.response,
			}

			req, err := http.NewRequest("GET", "/info/"+tt.args.urlPath, nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.GetHeroInfo)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.wantStatusCode)
			}
		})
	}
}
