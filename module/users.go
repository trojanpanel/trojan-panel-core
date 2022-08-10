package module

type Users struct {
	Id        *int    `ddb:"id"`
	AccountId *int    `ddb:"account_id"`
	ApiPort   *int    `ddb:"api_port"`
	Password  *string `ddb:"password"`
	Download  *int    `ddb:"download"`
	Upload    *int    `ddb:"upload"`
}
