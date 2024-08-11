package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

var server Server
var wg sync.WaitGroup

func startWebServer(ch chan struct{}, shutdownWG *sync.WaitGroup) {
	shutdownWG.Add(1)
	defer shutdownWG.Done()
	server = *NewServer()
	wg = sync.WaitGroup{}
	http.HandleFunc("/add", getAdd)
	http.HandleFunc("/remove", getRemove)

	go func() {
		err := http.ListenAndServe(":3333", nil)
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
	<-ch
	fmt.Println("Shutting down...")
	server.stopAllWorkers()
	fmt.Println("Waiting for all workers to stop")
	wg.Wait()

}

type AddBody struct {
	Token string `json:"token"`
	Uuid  string `json:"uuid"`
}

func getAdd(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /add request\n")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var reqBody AddBody
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received JSON: %+v\n", reqBody)
	response := fmt.Sprintf("Received data: Token = %s, UUID = %s", reqBody.Token, reqBody.Uuid)

	server.startWorker(reqBody.Uuid, reqBody.Token, &wg)
	io.WriteString(w, response)
}

type RemoveBody struct {
	UUID string `json:"uuid"`
}

func getRemove(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /remove request\n")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	var reqBody RemoveBody
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received JSON: %+v\n", reqBody)
	response := fmt.Sprintf("Received data: UUID = %s", reqBody.UUID)

	server.stopWorker(reqBody.UUID)
	io.WriteString(w, response)
}
