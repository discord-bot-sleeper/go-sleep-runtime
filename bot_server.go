package main

import (
	"fmt"
	"sync"
)

type Server struct {
	Workers map[string]*Worker
	mu      sync.Mutex
}

func NewServer() *Server {
	return &Server{
		Workers: make(map[string]*Worker),
	}
}

func (s *Server) startWorker(uuid string, token string, wg *sync.WaitGroup) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.Workers[uuid]; ok {
		fmt.Println("Bot with UUID " + uuid + " already exists")
		return
	}

	worker := &Worker{
		Token:    token,
		UUID:     uuid,
		StopChan: make(chan struct{}),
	}
	s.Workers[uuid] = worker
	wg.Add(1)
	go worker.Start(wg, &s.Workers)
}

func (s *Server) stopWorker(uuid string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if worker, exists := s.Workers[uuid]; exists {
		close(worker.StopChan)
		delete(s.Workers, uuid)
	} else {
		fmt.Println("Worker was not found for uuid ", uuid)
	}
}

func (s *Server) stopAllWorkers() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, v := range s.Workers {
		fmt.Println("Stopchan called LOL")
		close(v.StopChan)
	}
	fmt.Println("now clearing queue LOL")
	for k := range s.Workers {
		delete(s.Workers, k)
	}

}
