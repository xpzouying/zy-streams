package main

import "log"

type StdoutSink struct {
	in chan interface{}
}

func NewStdoutSink() *StdoutSink {
	s := &StdoutSink{
		in: make(chan interface{}),
	}

	go s.running()

	return s
}

func (s *StdoutSink) In() chan<- interface{} {
	return s.in
}

func (s *StdoutSink) running() {
	for elem := range s.in {
		log.Printf("sink: got element: %v", elem)
	}
}
