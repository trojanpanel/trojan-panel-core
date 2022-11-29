package dto

type NaiveProxyConfigDto struct {
	ApiPort uint
	Port    uint
	Ip      string
}

type NaiveProxyAddUserDto struct {
	Username string
	Pass     string
}
