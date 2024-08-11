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

func (w *Worker) Start(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Bot worker started with uuid = ", w.UUID)
	terminateChan := make(chan struct{})
	addBot(terminateChan, w.UUID, w.Token)

	for {
		select {
		case <-w.StopChan:
			fmt.Println("stopping bot with uuid ", w.UUID)
			close(terminateChan)
			return

		default:
			time.Sleep(500 * time.Millisecond)
		}

	}
}
