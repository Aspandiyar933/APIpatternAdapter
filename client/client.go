package client

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Aspandiyar933/APIpatternAdapter/pattern"
	"github.com/Aspandiyar933/APIpatternAdapter/store"
	"github.com/gorilla/mux"
)

type ClientService struct {
	store store.Store
}

func NewClientService(s store.Store) *ClientService {
	return &ClientService{
		store: s,
	}
}

func (s *ClientService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/do", s.HandleCreateTask).Methods("POST")
	r.HandleFunc("/do/{id}", s.HandleGetTask).Methods("GET")
}

func (s *ClientService) HandleCreateTask(w http.ResponseWriter, r *http.Request) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	var create *pattern.Todo
	err = json.Unmarshal(body, &create)
	if err != nil {
		return err
	}
	_, err = s.store.CreateTask(create)
	if err != nil {
		return err
	}
	return nil
}

func (s *ClientService) HandleGetTask(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		return nil 
	}
	_, err := s.store.GetTask(id)
	if err != nil {
		return nil
	}
	return nil 
}