package repository

import (
	"context"
	_ "embed"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

//go:embed lua/atomic_ticket.lua
var atomicTicketScript string

type QueueRepository struct {
	rdb *redis.Client
}

func NewQueueRepository(rdb *redis.Client) *QueueRepository {
	return &QueueRepository{rdb: rdb}
}

func (r *QueueRepository) TakeNextNumber(ctx context.Context, merchantID string) (int64, error) {
	// 1. Construct key: queue:merchant-123:2026-01-28
	date := time.Now().Format("2006-01-02")
	key := fmt.Sprintf("queue:%s:%s", merchantID, date)

	// 2. Constants
	ttl := 86400 // 24 Hours

	// 3. Execute Lua Script
	// KEYS[1] = key
	// ARGV[1] = ttl
	cmd := r.rdb.Eval(ctx, atomicTicketScript, []string{key}, ttl)

	if err := cmd.Err(); err != nil {
		return 0, fmt.Errorf("failed to execute lua script: %w", err)
	}

	return cmd.Int64()
}
