package scheduler

import "spider/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan chan chan engine.Request
}

func (q *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (q *QueuedScheduler) ReadyWorkChan(in chan engine.Request) {
	q.workerChan <- in
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) Run () {

	q.requestChan = make(chan engine.Request)
	q.workerChan = make(chan chan engine.Request)
	go func() {
		//创建一个request队列和一个worker队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest  engine.Request
			var activeWorker chan engine.Request

			//如果两个队列都有值的话就把两个队列的第一个给取出来
			if len(requestQ) >0 && len(workerQ)>0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <- q.requestChan:
				requestQ = append(requestQ, r)
			case w := <- q.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				//让request队列的第一个request交给worker队列的第一个worker处理
				//并把它们从队列中删除
				//此时activeWorker等于外面传进来的in
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
