package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/mugabwa/little-key-value/internal/storage"
)


type Server struct {
	storage storage.Storage
}

func New() *Server {
	return &Server{
		storage: storage.NewMapStorage(),
	}
}

func (s *Server) Server(addr string) error {
	http.HandleFunc("/kv/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			s.get(w, r)
		case http.MethodPut:
			s.set(w, r)
		case http.MethodDelete:
			s.delete(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	return http.ListenAndServe(addr, nil)
}

func (s *Server) set(w http.ResponseWriter, r *http.Request) {
	value, err := io.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Sprintf("Failed to read request body: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	if len(value) == 0 {
		http.Error(w, "Value is empty", http.StatusBadRequest)
		return
	}

	key := r.URL.Path[len("/kv/"):]
	if len(key) == 0 {
		http.Error(w, "Key is empty", http.StatusBadRequest)
		return
	}

	if err := s.storage.Set(key, string(value)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/kv/"):]
	value, err := s.storage.Get(key)
	if err != nil {
		if errors.Is(err, &storage.NotFoundError{}) {
			http.Error(w, "Key not found", http.StatusNotFound)
			return
		}
		msg := fmt.Sprintf("Failed to get value: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	w.Write([]byte(value))
	w.Write([]byte("\n"))
}

func (s *Server) delete(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/kv/"):]
	if err := s.storage.Delete(key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
