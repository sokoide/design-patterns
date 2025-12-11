package usecase

import (
	"abstract-factory-example/domain"
	"errors"
	"fmt"
)

type FurnishingService struct {
	factory domain.FurnitureFactory
}

func NewFurnishingService(f domain.FurnitureFactory) *FurnishingService {
	return &FurnishingService{factory: f}
}

func (s *FurnishingService) FurnishRoom() error {
	if s.factory == nil {
		return errors.New("furniture factory is not set")
	}
	chair := s.factory.CreateChair()
	sofa := s.factory.CreateSofa()

	fmt.Println(chair.SitOn())
	fmt.Println(sofa.LieOn())
	return nil
}
