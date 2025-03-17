package bo

type HysteriaUserTraffic struct {
	Tx int64 `json:"tx"` // upload
	Rx int64 `json:"rx"` // download
}

type HysteriaAuthBo struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
}
