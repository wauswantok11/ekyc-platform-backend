package services

import (
	"git.inet.co.th/ekyc-platform-backend/module/frontweb/ports"
	"git.inet.co.th/ekyc-platform-backend/module/frontweb/repositories"
)

type Service struct {
	repo ports.Repository
}

func New(repo *repositories.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
