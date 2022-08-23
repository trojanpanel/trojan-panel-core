package redis

import "github.com/gomodule/redigo/redis"

type zSetRds struct {
}

// map[score]member  添加元素
func (z *zSetRds) ZAdd(key string, mp map[interface{}]interface{}) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("zadd", redis.Args{}.Add(key).AddFlat(mp)...))
}

// 	增加元素权重
func (z *zSetRds) ZUncrBy(key string, increment, member interface{}) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("zuncrby", key, increment, member))
}

// 	增加元素权重
func (z *zSetRds) ZCard(key string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("zcard", key))
}

// 	返回指定元素的排名
func (z *zSetRds) ZEank(key string, member interface{}) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("zrank", key, member))
}

// 	返回指定元素的权重
func (z *zSetRds) ZScore(key string, member interface{}) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("zscore", key, member))
}

// 返回集合两个权重间的元素数
func (z *zSetRds) ZCount(key string, min, max interface{}) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("zcount", key, min, max))
}

// 返回指定区间内的元素
func (z *zSetRds) ZRange(key string, start, stop interface{}, withScore ...bool) *Reply {
	conn := pool.Get()
	defer conn.Close()
	if len(withScore) > 0 && withScore[0] {
		return getReply(conn.Do("zrange", key, start, stop, withScore))
	}
	return getReply(conn.Do("zrange", key, start, stop))
}

// 倒序返回指定区间内的元素
func (z *zSetRds) ZRevrange(key string, start, stop interface{}, withScore ...bool) *Reply {
	conn := pool.Get()
	defer conn.Close()
	if len(withScore) > 0 && withScore[0] {
		return getReply(conn.Do("zrevrange", key, start, stop, withScore))
	}
	return getReply(conn.Do("zrevrange", key, start, stop))
}
