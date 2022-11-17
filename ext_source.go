package main

var _ Source = (*ChanSource)(nil)

type ChanSource struct {
	in chan interface{}
}

func NewChanSource(in chan interface{}) *ChanSource {
	return &ChanSource{in}
}

func (s *ChanSource) Out() <-chan interface{} {
	return s.in
}

func (s *ChanSource) Via(flow Flow) Flow {
	go DoStream(s, flow)

	return flow
}
