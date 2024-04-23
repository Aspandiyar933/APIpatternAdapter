package api

import "github.com/Aspandiyar933/APIpatternAdapter/store"


type APIServer struct {
	addr string
	store store.Store
}