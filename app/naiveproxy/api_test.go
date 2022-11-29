package naiveproxy

import "testing"

func TestNaiveProxyListUsers(t *testing.T) {
	api := NewNaiveProxyApi(30883)
	users, err := api.ListUsers()
	if err != nil {
		println(err.Error())
	}
	for _, user := range *users {
		println(user.AuthUserDeprecated)
	}
}
