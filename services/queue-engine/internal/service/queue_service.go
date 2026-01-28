package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/herman-xphp/qtime-monorepo/services/queue-engine/internal/config"
	"github.com/herman-xphp/qtime-monorepo/services/queue-engine/internal/repository"
	"github.com/twmb/franz-go/pkg/kgo"
)

type QueueService struct {
	repo       *repository.QueueRepository
	kafka      *kgo.Client
	kafkaTopic string
}

type TicketCreatedEvent struct {
	TicketNumber int64     `json:"ticket_number"`
	MerchantID   string    `json:"merchant_id"`
	Timestamp    time.Time `json:"timestamp"`
}

func NewQueueService(repo *repository.QueueRepository, kafka *kgo.Client, cfg *config.Config) *QueueService {
	return &QueueService{
		repo:       repo,
		kafka:      kafka,
		kafkaTopic: cfg.Redpanda.Topic, // "queue-events"
	}
}

func (s *QueueService) GenerateTicket(ctx context.Context, merchantID string) (int64, error) {
	// 1. Get Atomic Ticket Number from Redis
	ticketNumber, err := s.repo.TakeNextNumber(ctx, merchantID)
	if err != nil {
		return 0, err
	}

	// 2. Publish Event to Redpanda (Async)
	go s.publishTicketCreated(merchantID, ticketNumber)

	return ticketNumber, nil
}

func (s *QueueService) publishTicketCreated(merchantID string, ticketNum int64) {
	event := TicketCreatedEvent{
		TicketNumber: ticketNum,
		MerchantID:   merchantID,
		Timestamp:    time.Now(),
	}

	payload, _ := json.Marshal(event)

	record := &kgo.Record{
		Topic: s.kafkaTopic,
		Key:   []byte(merchantID), // Key by MerchantID for ordering guarantees per merchant
		Value: payload,
	}

	// Fire and Forget (for now, or handle callback/error logging)
	s.kafka.Produce(context.Background(), record, func(r *kgo.Record, err error) {
		if err != nil {
			log.Printf("❌ Failed to publish event: %v", err)
		} else {
			log.Printf("✅ Event Published: %s", string(payload))
		}
	})
}
