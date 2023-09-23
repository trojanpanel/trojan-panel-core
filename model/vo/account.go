package vo

type AccountHysteriaVo struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}

type AccountVo struct {
	Id       uint     `json:"id"`
	Username string   `json:"username"`
	RoleId   uint     `json:"roleId"`
	Deleted  uint     `json:"deleted"`
	Roles    []string `json:"roles"`
}
