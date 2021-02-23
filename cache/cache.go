package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client
var ctx = context.TODO()
var expiration time.Duration = 5 * time.Second

// SetupDBConnection initializes Redis connection
func SetupDBConnection(addr, password string, DB int, cacheLifeSpan int64) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       DB,
	})

	expiration = time.Duration(cacheLifeSpan) * time.Second
}

/*
CachedFunctionCall tries to retrieve the value by `key`,
if the storage doesn't have the record, calls `targetFunction`
and uses *returned value* fill the storage with.
Returns the value of `targetFunction` call in this case.
*/
func CachedFunctionCall(key string, targetFunction func() (string, error)) (string, error) {
	result, err := rdb.Get(ctx, key).Result()

	if err == nil {
		return result, nil
	} else if err != redis.Nil {
		panic(err)
	}

	result, err = targetFunction()
	rdb.Set(ctx, key, result, expiration)

	return result, err
}
