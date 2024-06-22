package service

import (
	"dota2_info_service/internal/entity"
	"reflect"
	"testing"
)

func TestDotabuffService_GetHeroByName(t *testing.T) {
	type args struct {
		heroName string
	}
	tests := []struct {
		name    string
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DotabuffService{}
			got, err := s.GetHeroByName(tt.args.heroName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetHeroByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Name, tt.want.Name) {
				t.Errorf("GetHeroByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
