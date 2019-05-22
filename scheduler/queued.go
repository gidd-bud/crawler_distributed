package scheduler

import "IMOOC/crawler_distributed/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)
	go func() {
		var requsetQ []engine.Request
		var workerQ [] chan engine.Request
		for{
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requsetQ) > 0 && len(workerQ) > 0 {
				activeRequest = requsetQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <- s.requestChan:
				requsetQ = append(requsetQ, r)
			case w := <- s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requsetQ = requsetQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}


