package main

import (
	"log"
	"sync"
)

type MapFunction[T, R any] func(T) R

var _ Flow = (*Map[any, any])(nil)

type Map[T, R any] struct {
	mapFunction MapFunction[T, R]

	in  chan interface{}
	out chan interface{}
}

func NewMap[T, R any](mapFunction MapFunction[T, R]) *Map[T, R] {

	mapFlow := &Map[T, R]{
		mapFunction: mapFunction,

		in:  make(chan interface{}),
		out: make(chan interface{}),
	}

	go mapFlow.doStream()

	return mapFlow
}

func (m *Map[T, R]) In() chan<- interface{} {
	return m.in
}

func (m *Map[T, R]) Out() <-chan interface{} {
	return m.out
}

func (m *Map[T, R]) Via(flow Flow) Flow {

	go m.transmit(flow)

	return flow
}

func (m *Map[T, R]) To(sink Sink) {

	m.transmit(sink)

}

func (m *Map[T, R]) doStream() {

	wg := new(sync.WaitGroup)

	for elem := range m.in {

		wg.Add(1)
		go func(element T) {
			defer wg.Done()

			result := m.mapFunction(element)

			log.Printf("map.doStream: elem=%v result=%v", element, result)

			m.out <- result

		}(elem.(T))

	}

	wg.Wait()

	close(m.out)
}

func (m *Map[T, R]) transmit(inlet Inlet) {

	for elem := range m.Out() {
		inlet.In() <- elem
	}

	close(inlet.In())
}
