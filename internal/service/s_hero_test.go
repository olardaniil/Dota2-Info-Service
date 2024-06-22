package service

import (
	"dota2_info_service/internal/entity"
	"dota2_info_service/internal/repository"
	"dota2_info_service/internal/repository/redis"
	"dota2_info_service/pkg/database"
	"github.com/alicebob/miniredis/v2"
	"log"
	"reflect"
	"testing"
)

func TestHeroService_GetHeroByName(t *testing.T) {

	type fields struct {
		cacheHeroRepo repository.CacheHero
		dotabuff      Dotabuff
	}
	type args struct {
		heroName string
	}

	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	clientRedis, err := database.NewRedis(mr.Addr(), "", 0)
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Hero
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				heroName: "bristleback",
			},
			want: entity.Hero{
				Name: "bristleback",
			},
			wantErr: false,
			fields: fields{
				dotabuff: NewDotabuffService(),
				cacheHeroRepo: repository.Repository{
					CacheHero: redis.NewCacheHeroRepo(clientRedis),
				},
			},
		},
		{
			name: "Negative",
			args: args{
				heroName: "brist",
			},
			want: entity.Hero{
				Name: "",
			},
			wantErr: true,
			fields: fields{
				dotabuff: NewDotabuffService(),
				cacheHeroRepo: repository.Repository{
					CacheHero: redis.NewCacheHeroRepo(clientRedis),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &HeroService{
				cacheHeroRepo: tt.fields.cacheHeroRepo,
				dotabuff:      tt.fields.dotabuff,
			}
			got, err := s.GetHeroByName(tt.args.heroName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHeroByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Name, tt.want.Name) {
				log.Println()
				t.Errorf("GetHeroByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
