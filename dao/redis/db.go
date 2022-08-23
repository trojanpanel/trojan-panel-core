package redis

type dbRds struct {
}

func SelectDb(db int) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("select", db))
}
