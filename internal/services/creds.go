package services

import (
	"context"
	"github.com/RyanTrue/GophKeeper/internal/models"
	"github.com/RyanTrue/GophKeeper/internal/repository"
)

type CredsService struct {
	repo repository.CredsSecrets
}

var _ Creds = (*CredsService)(nil)

func NewCredsService(repo repository.CredsSecrets) Creds {
	return &CredsService{
		repo: repo,
	}
}

func (s *CredsService) GetList(ctx context.Context, userID int) ([]*models.CredsSecret, error) {
	return s.repo.GetList(ctx, userID)
}

func (s *CredsService) SetList(ctx context.Context, list []models.CredsSecret) error {
	return s.repo.SetList(ctx, list)
}
