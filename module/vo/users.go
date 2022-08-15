package vo

type UserApiVo struct {
	Password string `json:"password"`
	Download int    `json:"download"`
	Upload   int    `json:"upload"`
}
