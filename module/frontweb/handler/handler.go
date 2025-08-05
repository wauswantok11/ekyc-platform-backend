package handler

import (
	"git.inet.co.th/ekyc-platform-backend/module/frontweb/ports"
	"git.inet.co.th/ekyc-platform-backend/module/frontweb/services"
) 

type Handler struct {
	svc ports.Service
}

func NewHandler(svc *services.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}
