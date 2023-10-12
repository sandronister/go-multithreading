package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64

func main() {
	//m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		//number++
		//m.Unlock()
		atomic.AddUint64(&number, 1)
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante %d", number)))
	})
	http.ListenAndServe(":8000", nil)
}
