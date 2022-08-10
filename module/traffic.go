package module

import "time"

type Traffic struct {
	Id         *int       `ddb:"id"`
	UserId     *int       `ddb:"user_id"`
	ApiPort    *int       `ddb:"api_port"`
	Download   *int       `ddb:"download"`
	Upload     *int       `ddb:"upload"`
	CreateTime *time.Time `ddb:"create_time"`
	UpdateTime *time.Time `ddb:"update_time"`
}
