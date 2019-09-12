package utils

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
	"whacos/pkg/logging"
	"whacos/pkg/settings"
)

var (
	conn redis.Conn
	err  error
)

const (
	commandNameAuth   string = "AUTH"   // auth
	commandNamePing   string = "PING"   // ping
	commandNameSet    string = "SET"    // get
	commandNameGet    string = "GET"    // set
	commandNameExists string = "EXISTS" // exists
	commandNameDel    string = "DEL"    // delete
	commandNameSetnx  string = "SETNX"  // setnx
	commandNameExpire string = "EXPIRE" // expire
)

// 初始化
func init() {
	// 在初始化代码中连接redis
	pool := newPool()
	conn = pool.Get()

	//conn, err = redis.Dial("tcp", "148.70.20.37:6379")
	if err != nil {
		logging.Info("Connect to redis error , this error reason is ", err, err.Error())
	}

}

// 连接池
func newPool() *redis.Pool {
	return &redis.Pool{
		MaxActive:   settings.RedisConfig.MaxActive,
		MaxIdle:     settings.RedisConfig.MaxIdle,
		IdleTimeout: settings.RedisConfig.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(settings.RedisConfig.Network, settings.RedisConfig.Host)
			if err != nil {
				return nil, err
			}

			if _, err := conn.Do(commandNameAuth, settings.RedisConfig.Password); err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}

			_, err := conn.Do(commandNamePing)

			return err
		},
	}
}

// 关闭连接
func CloseRedis() {
	err = conn.Close()
	if err != nil {
		logging.Info("Close to redis error , this error reason is ", err, err.Error())
	}
	s1, e2 := redis.String(conn.Do(commandNameGet, "xing"))
	fmt.Println("ss水水水水", s1, e2)
}

// 设置过期时间
func SetExpire() {

}

func SetString(key string, value ...string) (reply string, err error) {
	reply, err = redis.String(conn.Do(commandNameSet, key, value))
	if err != nil {
		logging.Info("set key-value to redis error , this error reason is ", err, err.Error())
		return
	}
	return reply, nil
}

func GetString(key string) (value string, err error) {
	value, err = redis.String(conn.Do(commandNameGet, key))
	if err != nil {
		logging.Info("set key-value to redis error , this error reason is ", err, err.Error())
		return
	}
	return value, nil
}

func DeleteKey(key string) {

}
