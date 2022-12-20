package dto

type NaiveProxyConfigDto struct {
	ApiPort      uint
	Port         uint
	NodeServerIp string
}

type NaiveProxyAddUserDto struct {
	Username string
	Pass     string
}
