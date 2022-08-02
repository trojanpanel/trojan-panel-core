package module

type UsersXray struct {
	Password *string `ddb:"password"`
	Quota    *int    `ddb:"quota"`
	Download *int    `ddb:"download"`
	Upload   *int    `ddb:"upload"`
}
