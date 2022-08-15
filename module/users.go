package module

type Users struct {
	Id        *uint   `ddb:"id"`
	AccountId *uint   `ddb:"account_id"`
	ApiPort   *uint   `ddb:"api_port"`
	Password  *string `ddb:"password"`
	Download  *int    `ddb:"download"`
	Upload    *int    `ddb:"upload"`
}
