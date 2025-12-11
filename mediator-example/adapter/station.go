package adapter

import (
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
	// Free the platform for the next train.
	s.isPlatformFree = true
	if len(s.trainQueue) == 0 {
		s.lock.Unlock()
		return
	}
	// Dequeue next train and reserve the platform for it.
	firstTrain := s.trainQueue[0]
	s.trainQueue = s.trainQueue[1:]
	s.isPlatformFree = false
	s.lock.Unlock()

	// Call out to the train without holding the lock to avoid deadlocks.
	firstTrain.PermitArrival()
}
