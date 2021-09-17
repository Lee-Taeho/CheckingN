package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router interface {
	Serve(hostAndIpBinding string) error
	Stop() error
	AddRoute(methods, path string, handler func(w http.ResponseWriter, r *http.Request))
	GetRouter() *mux.Router
}
