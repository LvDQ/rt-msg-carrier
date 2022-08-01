package userinfo

import (
	"reflect"
	"testing"
)

func TestCreateNewUser(t *testing.T) {
	tests := []struct {
		name     string
		wantUser User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUser := CreateNewUser(); !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("CreateNewUser() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestGetUserDetailById(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name     string
		args     args
		wantUser User
	}{
		// TODO: Add test cases.
		{
			name:     "test1", // this is useful only when location wrong test cases.
			args:     args{id: 2},
			wantUser: User{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUser := GetUserDetailById(tt.args.id); !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("GetUserDetailById() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}
