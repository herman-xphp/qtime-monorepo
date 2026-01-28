package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/herman-xphp/qtime-monorepo/services/queue-engine/internal/repository"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestTakeNextNumber(t *testing.T) {
	// 1. Setup Miniredis (Mock Redis Server)
	s := miniredis.RunT(t)
	defer s.Close()

	// 2. Connect go-redis to Miniredis
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	// 3. Init Repository
	repo := repository.NewQueueRepository(rdb)
	ctx := context.Background()
	merchantID := "TokoTest"

	// --- Case A: First Ticket (Must be 1 and set TTL) ---
	ticket1, err := repo.TakeNextNumber(ctx, merchantID)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), ticket1)

	// Check Key
	date := time.Now().Format("2006-01-02")
	key := "queue:" + merchantID + ":" + date

	// Check Value in Mock Redis
	val, _ := s.Get(key)
	assert.Equal(t, "1", val)

	// Check TTL (Must be 24 hours / 86400 seconds)
	ttl := s.TTL(key)
	assert.True(t, ttl > 0)

	// --- Case B: Second Ticket (Must be 2) ---
	ticket2, err := repo.TakeNextNumber(ctx, merchantID)
	assert.NoError(t, err)
	assert.Equal(t, int64(2), ticket2)
}
