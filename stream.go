package main

type Inlet interface {
	In() chan<- interface{}
}

type Outlet interface {
	Out() <-chan interface{}
}

type Viaer interface {
	Via(Flow) Flow
}

type Source interface {
	Outlet

	Viaer
}

type Flow interface {
	Inlet
	Outlet

	Via(Flow) Flow
	To(Sink)
}

type Sink interface {
	Inlet
}

// DoStream to process from src to dest.
func DoStream(src Outlet, dest Inlet) {
	for elem := range src.Out() {
		dest.In() <- elem
	}

	close(dest.In())
}
