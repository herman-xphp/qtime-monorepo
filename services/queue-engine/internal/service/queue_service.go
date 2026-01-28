package service

import (
	"context"

	"github.com/herman-xphp/qtime-monorepo/services/queue-engine/internal/repository"
)

type QueueService struct {
	repo *repository.QueueRepository
}

func NewQueueService(repo *repository.QueueRepository) *QueueService {
	return &QueueService{repo: repo}
}

func (s *QueueService) GenerateTicket(ctx context.Context, merchantID string) (int64, error) {
	return s.repo.TakeNextNumber(ctx, merchantID)
}
