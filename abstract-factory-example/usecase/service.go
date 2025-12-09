package usecase

import (
	"abstract-factory-example/domain"
	"fmt"
)

type FurnishingService struct {
	factory domain.FurnitureFactory
}

func NewFurnishingService(f domain.FurnitureFactory) *FurnishingService {
	return &FurnishingService{factory: f}
}

func (s *FurnishingService) FurnishRoom() {
	chair := s.factory.CreateChair()
	sofa := s.factory.CreateSofa()

	fmt.Println(chair.SitOn())
	fmt.Println(sofa.LieOn())
}
