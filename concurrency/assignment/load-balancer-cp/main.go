package main

import (
	"fmt"
	"sync"
)

const nWorkers = 3
const nRequesters = 5

type Work struct {
	data string
}

var mu = &sync.Mutex{}

var data = map[string]bool{}

func worker(in <-chan *Work, out chan<- *Work, number int) {
	for w := range in {
		w.data = fmt.Sprintf("worker %d accepted %s", number, w.data)
		out <- w

	}
}

func createRequest(in chan<- *Work, number int) {
	for {
		// TODO: answer here
		in <- &Work{data: fmt.Sprintf("request from client %d", number)}

	}
}

func receiver(out <-chan *Work) {
	for {
		//pada bagian ini masukkan data dari channel out
		//ke dalam map `data`
		//gunakan mutex untuk mengamankan penulisan ke map secara concurrent

		for {
			w := <-out
			mu.Lock()
			data[w.data] = true
			mu.Unlock()
		}
		// TODO: answer here
	}
}
