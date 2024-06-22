package redis

import (
	"dota2_info_service/internal/entity"
	"dota2_info_service/pkg/database"
	"github.com/alicebob/miniredis/v2"
	"reflect"
	"testing"
)

func TestCacheHeroRepo_GetHeroByName(t *testing.T) {
	type fields struct {
		redis *database.Redis
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
			want:    entity.Hero{},
			wantErr: true,
			fields: fields{
				redis: clientRedis,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CacheHeroRepo{
				redis: tt.fields.redis,
			}
			got, err := c.GetHeroByName(tt.args.heroName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHeroByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHeroByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCacheHeroRepo_SetHero(t *testing.T) {
	type fields struct {
		redis *database.Redis
	}
	type args struct {
		hero entity.Hero
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
		wantErr bool
	}{
		{
			name: "Positive",
			args: args{
				hero: entity.Hero{
					Name: "bristleback",
				},
			},
			wantErr: false,
			fields: fields{
				redis: clientRedis,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CacheHeroRepo{
				redis: tt.fields.redis,
			}
			if err := c.SetHero(tt.args.hero); (err != nil) != tt.wantErr {
				t.Errorf("SetHero() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
