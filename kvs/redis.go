package kvs

import (
	"os"

	"golang.org/x/xerrors"

	"github.com/gomodule/redigo/redis"
)

var Conn redis.Conn

func RunRedis() error {
	addr := os.Getenv("REDIS_ADDRESS") + ":" + os.Getenv("REDIS_PORT")
	var err error
	Conn, err = redis.Dial("tcp", addr)
	if err != nil {
		return err
	}
	return nil
}

func Set(key, value string, c redis.Conn) error {
	_, err := c.Do("SET", key, value)
	if err != nil {
		return xerrors.Errorf("Do error: %w", err)
	}
	return nil
}

func Get(key string, c redis.Conn) (string, error) {
	res, err := redis.String(c.Do("GET", key))
	if err != nil {
		return "", xerrors.Errorf("Do error: %w", err)
	}
	return string(res), nil
}
