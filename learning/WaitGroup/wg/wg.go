package wg

import (
	"sync"
)

type Wg struct {
	wg sync.WaitGroup
	ch chan struct{}
}

func (w *Wg) Add() {
	w.ch <- struct{}{}
	w.wg.Add(1)
}

func (w *Wg) Done() {
	<-w.ch
	w.wg.Done()
}

func (w *Wg) Wait(){
	w.wg.Wait()
}

func (w *Wg) Square(res chan interface{}) chan interface{}{
	out := make(chan interface{})
	s := make([]interface{}, 1)
	go func() {
		defer close(out)
		i := 0
		for r := range res{
			s[i] = r
			s = append(s, r)
			i++
		}
		out <- s
	}()

	return out
}

func NewWg(count int) *Wg {
	wg := sync.WaitGroup{}
	return &Wg{
		wg: wg,
		ch: make(chan struct{}, count),
	}
}
