package main

type workerProvider interface {
	do()
}

type azuerProvider string

func (w azuerProvider) do() {
	print("im on azuer")
}

type awsProvider string

func (w awsProvider) do() {
	print("im on aws")
}

type worker struct {
	concreteWorker workerProvider
}

func (ww worker) execute() {
	ww.concreteWorker.do()
}

func main() {
	worker := new(worker)
	worker.concreteWorker = new(awsProvider)
	worker.execute()

	//
	worker.concreteWorker = new(azuerProvider)
	worker.execute()
}
