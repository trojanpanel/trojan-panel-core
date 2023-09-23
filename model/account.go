package model

type Account struct {
	Id                 *uint   `ddb:"id"`
	Username           *string `ddb:"username"`
	Pass               *string `ddb:"pass"`
	Hash               *string `ddb:"hash"`
	Quota              *int    `ddb:"quota"`
	Download           *int    `ddb:"download"`
	Upload             *int    `ddb:"upload"`
	IpLimit            *int    `ddb:"ip_limit"`
	DownloadSpeedLimit *int    `ddb:"download_speed_limit"`
	UploadSpeedLimit   *int    `ddb:"upload_speed_limit"`
}
