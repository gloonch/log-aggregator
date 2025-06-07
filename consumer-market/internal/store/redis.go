package store

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(addr string) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	return &RedisStore{client: rdb}
}

func (r RedisStore) SaveCandle(symbol, timeFrame string, value []byte, ts int64) error {
	key := "market" + symbol + ":" + timeFrame
	return r.client.ZAdd(ctx, key, redis.Z{
		Score:  float64(ts),
		Member: value,
	}).Err()
}

func (r *RedisStore) RangeByScore(ctx context.Context, key string, start, end int64) ([]redis.Z, error) {
	return r.client.ZRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{
		Min: fmt.Sprintf("%d", start),
		Max: fmt.Sprintf("%d", end),
	}).Result()
}
