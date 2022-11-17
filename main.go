package main

import (
	"log"
	"strconv"
	"time"
)

func main() {

	source := NewChanSource(tickerChan(10, time.Millisecond*500))
	mapFlowItoa := NewMap(ConvertToString)
	mapFlowAppend := NewMap(AppendWord)
	sink := NewStdoutSink()

	source.
		Via(mapFlowItoa).
		Via(mapFlowAppend).
		To(sink)

}

var ConvertToString = func(i int) string {

	return strconv.Itoa(i)
}

var AppendWord = func(s string) string {

	return "HelloWorld_" + s
}

func tickerChan(count int, dur time.Duration) chan interface{} {
	ticker := time.NewTicker(dur)

	c := make(chan interface{})

	go func() {

		for i := 0; i < count; i++ {
			<-ticker.C

			c <- i
		}

		ticker.Stop()
		log.Printf("send all data into source channel, close channel.")
		close(c)
	}()

	return c
}
