package models

import (
	"context"
	"testing"
)

func tearDown(store store, ctx context.Context) {
	store.db.FlushAll(ctx)
}

func TestInitGetCounter(t *testing.T) {
	ctx := context.Background()
	defer tearDown(redisStore, ctx)
	counter := &Counter{}
	counter.Get(ctx)
	if counter.Count != 0 {
		t.Error("initailly counter must be zero but is ", counter.Count)
	}
}

func TestIncrCounter(t *testing.T) {
	ctx := context.Background()
	defer tearDown(redisStore, ctx)
	counter := &Counter{}
	counter.Incr(ctx)
	if counter.Count != 1 {
		t.Error("initailly counter must be zero but is ", counter.Count)
	}
}
