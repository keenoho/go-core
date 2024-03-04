package module

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/keenoho/go-core"
)

var storeRedis *redis.Pool

type RedisModule struct {
	core.Module
}

func (m *RedisModule) Init(app *core.App) {
	m.initRedis()
}

func (m *RedisModule) initRedis() {
	if storeRedis != nil {
		return
	}
	database := core.ConfigGet("REDIS_DATABASE")
	password := core.ConfigGet("REDIS_PASSWORD")
	host := core.ConfigGet("REDIS_HOST")
	port := core.ConfigGet("REDIS_PORT")

	address := fmt.Sprintf("%s:%s", host, port)
	db, _ := strconv.Atoi(database)

	storeRedis = &redis.Pool{
		MaxIdle:   2,    // idle connect num
		MaxActive: 5000, // max connect num
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address, redis.DialPassword(password), redis.DialDatabase(db), redis.DialWriteTimeout(10*time.Second), redis.DialReadTimeout(10*time.Second))
		},
		IdleTimeout: 240 * time.Second,
		Wait:        true,
	}

	m.testConnect()
}

func (m *RedisModule) Redis() *redis.Pool {
	return storeRedis
}

func (m *RedisModule) RedisConnect() redis.Conn {
	conn := storeRedis.Get()
	return conn
}

func (m *RedisModule) RedisExec(cmd string, params ...any) (reply interface{}, err error) {
	conn := m.RedisConnect()
	reply, err = conn.Do(cmd, params...)
	conn.Close()
	return reply, err
}

func (m *RedisModule) RedisGet(key string) (reply interface{}, err error) {
	return m.RedisExec("get", key)
}

func (m *RedisModule) RedisSet(key string, value any, ttl ...int) (reply interface{}, err error) {
	if len(ttl) > 0 {
		return m.RedisExec("set", key, value, "EX", ttl[0])
	} else {
		return m.RedisExec("set", key, value)
	}
}

func (m *RedisModule) RedisDelete(key string) (reply interface{}, err error) {
	return m.RedisExec("del", key)
}

func (m *RedisModule) testConnect() {
	conn := m.RedisConnect()
	res, err := conn.Do("ping")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("redis test: ping =", res, "ok")
}
