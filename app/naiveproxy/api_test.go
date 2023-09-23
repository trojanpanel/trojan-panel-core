package naiveproxy

import (
	"fmt"
	"testing"
	"trojan-panel-core/model/dto"
)

func TestNaiveProxyListUsers(t *testing.T) {
	api := NewNaiveProxyApi(30883)
	users, err := api.ListUsers()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	for _, user := range *users {
		fmt.Println(user.AuthUserDeprecated)
		fmt.Println(user.AuthPassDeprecated)
	}
}

func TestNaiveProxyGetUser(t *testing.T) {
	api := NewNaiveProxyApi(30883)
	user, index, err := api.GetUser("111111")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	if user != nil {
		fmt.Println(user.AuthUserDeprecated)
		fmt.Println(*index)
	}
}

func TestNaiveProxyAddUser(t *testing.T) {
	api := NewNaiveProxyApi(30883)
	userDto := dto.NaiveProxyAddUserDto{
		Username: "111111",
		Pass:     "111111",
	}
	if err := api.AddUser(userDto); err != nil {
		fmt.Printf("%v\n", err)
	}
}
func TestNaiveProxyDeleteUser(t *testing.T) {
	api := NewNaiveProxyApi(30883)
	if err := api.DeleteUser("111111"); err != nil {
		fmt.Printf("%v\n", err)
	}
}
