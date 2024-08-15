package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	Token    string
	UUID     string
	StopChan chan struct{}
}

func (w *Worker) Start(wg *sync.WaitGroup, workers *map[string]*Worker) {
	fmt.Println("Bot worker started with uuid = ", w.UUID)
	terminateChan := make(chan struct{})
	go addBot(terminateChan, w.UUID, w.Token, wg)

	for {
		select {
		case <-w.StopChan:
			fmt.Println("stopping bot with uuid ", w.UUID)
			close(terminateChan)
			return
		case <-terminateChan:
			delete(*workers, w.UUID)
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}

	}
}
