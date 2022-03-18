package redisx

import (
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func GetAndJsonToObject(redis *redis.Redis, key string, res interface{}) error {
	cartString, err := redis.Get(key)

	if err != nil {
		logx.Error(err)
		return err
	}

	if len(cartString) <= 0 {
		return nil
	}

	err = json.Unmarshal([]byte(cartString), res)
	if err != nil {
		logx.Error(err)
		return err
	}

	return nil
}
