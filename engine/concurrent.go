package engine

type ConcurrentEngine struct {
}
type Scheduler interface {
	Submit(Request)
}

func (e ConcurrentEngine) Run(seeds ...Request) {

}
