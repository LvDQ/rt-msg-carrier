package plat_user

import (
	"testing"
)

func TestLogin_Logut(t *testing.T) {
	user := NewPlatUser(123)
	user.Login()
	status, err := user.Is_Logged_In()
	if err != nil {
		t.Errorf("user not correct logged in ")
	}
	if status == true {
		t.Log("user logged in correctly")
	}
	user.Logout()
	status, err = user.Is_Logged_In()
	if err != nil {
		t.Errorf("user not correct logged out ")
	}
	if status == false {
		t.Log("user logged out")
	}
}
