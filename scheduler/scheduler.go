package scheduler

import "amapdistricts/types"

type SimpleScheduler struct {
	workerChan chan types.QueryCondition
}

func (s *SimpleScheduler) WorkChan() chan types.QueryCondition {
	return s.workerChan
}

func (s *SimpleScheduler) WorkReady(chan types.QueryCondition) {

}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan types.QueryCondition)

}

func (s *SimpleScheduler) Submit(r types.QueryCondition) {
	go func() { s.workerChan <- r }()
}
