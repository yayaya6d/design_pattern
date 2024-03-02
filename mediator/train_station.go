package mediator

type TrainStation struct {
	trainQueue     []Train
	isPlatformFree bool
}

func NewTrainStation() Mediator {
	return &TrainStation{
		isPlatformFree: true,
	}
}

func (s *TrainStation) CanArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}

	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *TrainStation) NotifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}

	if len(s.trainQueue) > 0 {
		// let first train arrive to platform.
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}

func (s *TrainStation) WaitingTrainCount() int {
	return len(s.trainQueue)
}
