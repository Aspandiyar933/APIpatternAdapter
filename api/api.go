package api

import (
	"github.com/Aspandiyar933/APIpatternAdapter/client"
	"github.com/Aspandiyar933/APIpatternAdapter/store"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	store store.Store
}

func NewAPIServer(addr string, store store.Store) *APIServer {
	return &APIServer{
		addr: addr,
		store: store,
	}
}

func (a *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/adapter/v1").Subrouter()

	clientService := client.NewClientService(a.store)
	clientService.RegisterRoutes(subrouter)

}