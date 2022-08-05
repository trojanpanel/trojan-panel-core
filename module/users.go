package module

import "time"

type UsersXray struct {
	Id         *int       `ddb:"id"`
	Username   *string    `ddb:"username"`
	Password   *string    `ddb:"password"`
	Quota      *int       `ddb:"quota"`
	Download   *int       `ddb:"download"`
	Upload     *int       `ddb:"upload"`
	CreateTime *time.Time `ddb:"create_time"`
	UpdateTime *time.Time `ddb:"update_time"`
}
