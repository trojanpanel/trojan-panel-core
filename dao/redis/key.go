package redis

import "github.com/gomodule/redigo/redis"

type keyRds struct {
}

// 	查找键 [*模糊查找]
func (k *keyRds) Keys(key string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("keys", key))
}

// 	判断key是否存在
func (k *keyRds) Exists(key string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("exists", key))
}

// 随机返回一个key
func (k *keyRds) RandomKey() *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("randomkey"))
}

// 返回值类型
func (k *keyRds) Type(key string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("type", key))
}

// 删除key
func (k *keyRds) Del(keys ...string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("del", redis.Args{}.AddFlat(keys)...))
}

//重命名
func (k *keyRds) Rename(key, newKey string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("rename", key, newKey))
}

// 仅当newkey不存在时重命名
func (k *keyRds) RenameNX(key, newKey string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("renamenx", key, newKey))
}

//	序列化key
func (k *keyRds) Dump(key string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("dump", key))
}

// 反序列化
func (k *keyRds) Restore(key string, ttl, serializedValue interface{}) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("restore", key, ttl, serializedValue))
}

// 秒
func (k *keyRds) Expire(key string, seconds int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("expire", key, seconds))
}

// 秒
func (k *keyRds) ExpireAt(key string, timestamp int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("expireat", key, timestamp))
}

// 毫秒
func (k *keyRds) Persist(key string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("persist", key))
}

// 毫秒
func (k *keyRds) PersistAt(key string, milliSeconds int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("persistat", key, milliSeconds))
}

// 秒
func (k *keyRds) TTL(key string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("ttl", key))
}

// 毫秒
func (k *keyRds) PTTL(key string) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("pttl", key))
}

//	同实例不同库间的键移动
func (k *keyRds) Move(key string, db int64) *Reply {
	conn := pool.Get()
	defer conn.Close()
	return getReply(conn.Do("move", key, db))
}
