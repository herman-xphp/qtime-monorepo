package config

import (
	"context"
	"log"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

func NewRedPandaClient(cfg *Config) *kgo.Client {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.Redpanda.Brokers...),
		kgo.AllowAutoTopicCreation(), // Auto create topic if not exists (Dev Mode)
	)
	if err != nil {
		log.Fatalf("❌ Failed to init Redpanda client: %v", err)
	}

	// Test Connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx); err != nil {
		log.Fatalf("❌ Failed to connect to Redpanda: %v", err)
	}

	log.Println("✅ Connected to Redpanda")
	return client
}
