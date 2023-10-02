package api

import (
	"encoding/json"
	"learning-go/storage"
	"learning-go/utils"
	"net/http"
)

type Server struct {
	listenAddr string
	store storage.Storage //from the storage interface
}

func NewServer(listenAddr string, store storage.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store: store,
	}
}

//  function for http handlers or urls

func (s *Server) Start() error {
	http.HandleFunc(("/users"), s.handleUserByID)
	http.HandleFunc("/users/id", s.handleDeleteUser)
	http.HandleFunc("/api", s.handleApi)
	http.HandleFunc("/api/requsts", s.handleApiRequests)
	http.HandleFunc("/api/all", s.handleApiAll)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleUserByID(w http.ResponseWriter, r *http.Request){
	user := s.store.Get(10)
	json.NewEncoder(w).Encode(user)
}

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request){
	user := s.store.Get(10)
	_ = utils.Round2Dec(10.345678)
	json.NewEncoder(w).Encode(user)
}