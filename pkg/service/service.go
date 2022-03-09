package service

import (
	"fmt"

	"github.com/romanyx/stack"
)

type Service interface {
	WorldOriginal() (string, error)
	WorldStacktrace() (string, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) WorldOriginal() (string, error) {
	return "", fmt.Errorf("hello debug world")
}

func (s *service) WorldStacktrace() (string, error) {
	return "", stack.Errorf("hello debug world")
}
