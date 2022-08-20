package module

type NodeXray struct {
	Id           *uint   `ddb:"id"`
	Protocol     *string `ddb:"protocol"`
	SSMethod     *string `ddb:"ss_method"`
	VlessId      *string `ddb:"vless_id"`
	VmessId      *string `ddb:"vmess_id"`
	VmessAlterId *string `ddb:"vmess_alter_id"`
}
