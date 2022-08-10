package module

type Account struct {
	Id                 *int    `ddb:"id"`
	Username           *string `ddb:"username"`
	Pass               *string `ddb:"pass"`
	Quota              *int    `ddb:"quota"`
	Download           *int    `ddb:"download"`
	Upload             *int    `ddb:"upload"`
	IpLimit            *int    `ddb:"ip_limit"`
	DownloadSpeedLimit *int    `ddb:"download_speed_limit"`
	UploadSpeedLimit   *int    `ddb:"upload_speed_limit"`
}
