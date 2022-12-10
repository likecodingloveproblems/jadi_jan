package models

import (
	"context"
	"strconv"
)

type Counter struct {
	Count uint64
}

func (c *Counter) Incr(ctx context.Context) {
	count := redisStore.db.Incr(ctx, "counter").Val()
	c.Count = uint64(count)
}

func (c *Counter) Get(ctx context.Context) {
	count := redisStore.db.Get(ctx, "counter").Val()
	count_int, _ := strconv.Atoi(count)
	c.Count = uint64(count_int)
}
