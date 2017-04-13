package gopool

import (
	"reflect"
	"sync"
)

type Worker struct {
	f          reflect.Value
	paramsList [][]reflect.Value
	wg         sync.WaitGroup
	Count      int
	MaxPool    int
	paramsChan chan []reflect.Value
	pool       chan int
}

func (this *Worker) Init(f interface{}, maxPool int) {
	this.f = reflect.ValueOf(f)
	this.MaxPool = maxPool
	this.paramsChan = make(chan []reflect.Value)
	this.pool = make(chan int, maxPool)
}

func (this *Worker) Push(params ...interface{}) {
	paramList := make([]reflect.Value, 0)
	for _, p := range params {
		paramList = append(paramList, reflect.ValueOf(p))
	}
	this.paramsList = append(this.paramsList, paramList)
	this.Count++
}

func (this *Worker) run() {
	for params := range this.paramsChan {
		this.wg.Add(1)
		go func(params []reflect.Value) {
			defer this.wg.Done()
			this.f.Call(params)
			<-this.pool
		}(params)
	}
}

func (this *Worker) Start() {
	go this.run()
	for _, params := range this.paramsList {
		this.pool <- 1
		this.paramsChan <- params
	}
	close(this.paramsChan)
}

func (this *Worker) Wait() {
	this.wg.Wait()
}
