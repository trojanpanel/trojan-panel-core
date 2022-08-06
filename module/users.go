package module

type Users struct {
	Id       *int    `ddb:"id"`
	Password *string `ddb:"password"`
	Quota    *int    `ddb:"quota"`
	Download *int    `ddb:"download"`
	Upload   *int    `ddb:"upload"`
}
