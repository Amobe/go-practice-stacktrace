package handler

import (
	"fmt"

	"github.com/amobe/go-practice-stacktrace/pkg/service"
	"github.com/romanyx/stack"
)

type Handler interface {
	HelloOriginal() (string, error)
	HelloStacktrace() (string, error)
}

type handler struct {
	service service.Service
}

func NewHandler(s service.Service) Handler {
	return &handler{
		service: s,
	}
}

func (h *handler) HelloOriginal() (string, error) {
	res, err := h.service.WorldOriginal()
	if err != nil {
		return "", fmt.Errorf("service world failed: %w", err)
	}
	return res, nil
}

func (h *handler) HelloStacktrace() (string, error) {
	res, err := h.service.WorldStacktrace()
	if err != nil {
		return "", stack.Errorf("service world failed: %w", err)
	}
	return res, nil
}
