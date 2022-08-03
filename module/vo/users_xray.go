package vo

type UsersXrayVo struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Quota    int    `json:"quota"`
	Download int    `json:"download"`
	Upload   int    `json:"upload"`
}
