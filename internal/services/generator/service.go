package generator

import uuid "github.com/satori/go.uuid"

type Service struct {
}

func (s *Service) UniqueID() string {
	return uuid.NewV4().String()
}

func NewGenerator() *Service {
	return new(Service)
}
