package userinfo

import (
	"reflect"
	"testing"
)

func TestGetFavoriteDetailsByUserid(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name         string
		args         args
		wantFavorate Favorite
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFavorate := GetFavoriteDetailsByUserid(tt.args.id); !reflect.DeepEqual(gotFavorate, tt.wantFavorate) {
				t.Errorf("GetFavoriteDetailsByUserid() = %v, want %v", gotFavorate, tt.wantFavorate)
			}
		})
	}
}
