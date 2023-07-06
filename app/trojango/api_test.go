package trojango

import (
	"fmt"
	"testing"
)

func TestTrojanGoListUsers(t *testing.T) {
	api := NewTrojanGoApi(30452)
	users, err := api.ListUsers()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	for _, user := range users {
		fmt.Println(user.GetUser().GetHash())
		fmt.Println(user.GetTrafficTotal())
	}
}
