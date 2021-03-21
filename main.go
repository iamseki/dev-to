package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/iamseki/dev-to/domain"
	"github.com/iamseki/dev-to/infra/repository"
	"github.com/iamseki/dev-to/usecases"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
			return
		}

		repo := repository.NewInMemoryRepository()
		u := usecases.NewFindEventInMemory(repo)

		events, err := u.Find(domain.Filter{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			body, _ := json.Marshal(events)
			w.Header().Set("Content-Type", "application/json")
			w.Write(body)
		}
	})

	addr := `:8080`
	srv := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Println(`listen on localhost`, addr)
	log.Fatalln(srv.ListenAndServe())
}
