package dto

type NaiveProxyConfigDto struct {
	ApiPort uint
	Port    uint
	Domain  string
}

type NaiveProxyAddUserDto struct {
	Username string
	Pass     string
}
