package redis

import "github.com/gomodule/redigo/redis"

type bitRds struct {
}

func (b *bitRds) SetBit(key string, offset, value int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("setbit", key, offset, value))
}

func (b *bitRds) GetBit(key string, offset int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("getbit", key, offset))
}

func (b *bitRds) BitCount(key string, interval ...int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	if len(interval) == 2 {
		return getReply(conn.Do("bitcount", key, interval[0], interval[1]))
	}
	return getReply(conn.Do("bitcount", key))
}

// opt 包含 and、or、xor、not
func (b *bitRds) BitTop(opt, destKey string, keys ...string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("bitop", opt, redis.Args{}.Add(keys).AddFlat(keys)))
}
