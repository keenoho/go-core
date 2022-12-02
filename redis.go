package core

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

var Redis redis.Pool

func LoadRedis() {
	conf := GetConfig()
	address := fmt.Sprintf("%s:%s", conf.RedisHost, conf.RedisPort)
	db, _ := strconv.Atoi(conf.RedisDatabase)

	Redis = redis.Pool{
		MaxIdle:   2,    // 空闲链接数
		MaxActive: 5000, // 最大链接数
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address, redis.DialPassword(conf.RedisPassword), redis.DialDatabase(db), redis.DialWriteTimeout(10*time.Second), redis.DialReadTimeout(10*time.Second))
		},
		IdleTimeout: 240 * time.Second,
		Wait:        true,
	}
	testRedisConnect()
}

func testRedisConnect() {
	conn := Redis.Get()
	res, err := conn.Do("ping")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("redis test: ping =", res, "ok")
}

func GetRedisConnect() redis.Conn {
	conn := Redis.Get()
	return conn
}

func RedisExec(params ...any) (reply interface{}, err error) {
	conn := Redis.Get()
	reply, err = conn.Do("set", params...)
	conn.Close()
	return reply, err
}

func RedisSet(key string, value any, ttl ...int) (reply interface{}, err error) {
	conn := Redis.Get()
	if len(ttl) > 0 {
		reply, err = conn.Do("set", key, value, "EX", ttl[0])
	} else {
		reply, err = conn.Do("set", key, value)
	}
	conn.Close()
	return reply, err
}

func RedisGet(key string) (reply interface{}, err error) {
	conn := Redis.Get()
	reply, err = conn.Do("get", key)
	conn.Close()
	return reply, err
}
