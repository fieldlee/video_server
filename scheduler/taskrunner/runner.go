package taskrunner

import "log"

type Runner struct {
	Controller controlChan
	Error controlChan
	Data  dataChan
	dataSize int
	longLive bool
	Dispatcher fn
	Executer fn
}

func NewRunner(size int,longlive bool,d fn,e fn) *Runner {
	return &Runner{
		Controller:make(chan string,1),
		Error: make(chan string,1),
		dataSize:size,
		Data: make(chan interface{} ,size),
		longLive:longlive,
		Dispatcher:d,
		Executer:e,
	}
}

func (r *Runner)startDispatcher()  {
	defer func() {
		if !r.longLive{
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()

	for{
		select {
		case c := <- r.Controller:
			if c == READY_TO_DISPATCH{
				err := r.Dispatcher(r.Data)
				if err != nil {
					r.Error <- CLOSE
				}else{
					r.Controller <- READY_TO_EXECUTE
				}
			}

			if c == READY_TO_EXECUTE{
				err := r.Executer(r.Data)
				if err != nil {
					r.Error <- CLOSE
				}else{
					r.Controller <- READY_TO_DISPATCH
				}
			}

		case e := <- r.Error:
			if e == CLOSE {
				return
			}
		default:
			log.Printf("runner default")
		}
	}
}

func (r *Runner)StartAll()  {
	r.Controller <- READY_TO_DISPATCH
	r.startDispatcher()
}