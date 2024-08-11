package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func main() {

	ShutDown := make(chan struct{})
	shutdownWG := sync.WaitGroup{}
	go monitor()
	go startWebServer(ShutDown, &shutdownWG)

	os_shutdown := make(chan os.Signal, 1)
	signal.Notify(os_shutdown, syscall.SIGINT, syscall.SIGTERM)

	<-os_shutdown
	fmt.Println("OS shutdown...")
	close(ShutDown)
	shutdownWG.Wait()

}

func monitor() {
	var m runtime.MemStats
	for {
		runtime.ReadMemStats(&m)

		fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
		fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
		fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
		fmt.Printf("\tNumGC = %v\n", m.NumGC)

		time.Sleep(1 * time.Second)
	}
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
