package adapter

import (
	"fmt"
	"mediator-example/domain"
	"sync"
)

type StationManager struct {
	isPlatformFree bool
	lock           *sync.Mutex
	trainQueue     []domain.Train
}

func NewStationManager() *StationManager {
	return &StationManager{
		isPlatformFree: true,
		lock:           &sync.Mutex{},
	}
}

func (s *StationManager) CanArrive(t domain.Train) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) NotifyAboutDeparture() {
	s.lock.Lock()
	defer s.lock.Unlock()
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrain := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrain.PermitArrival()
	}
}
