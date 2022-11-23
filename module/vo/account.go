package vo

type AccountHysteriaVo struct {
	Id uint `json:"id"`
}

type AccountVo struct {
	Id       uint     `json:"id"`
	Username string   `json:"username"`
	RoleId   uint     `json:"roleId"`
	Deleted  uint     `json:"deleted"`
	Roles    []string `json:"roles"`
}

type AccountAuthVo struct {
	Username string `json:"username"`
	Pass     string `json:"pass"`
}
