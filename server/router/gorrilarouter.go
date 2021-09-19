package router

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// golang classes are made from structs
// for example, this struct will be the instance variables of the GorillaRouter class
// names that start with a lowercase letter are equivalent to "private" in Java (the instance variable router starts with a lowercase letter)
// names with uppercase starting letter are public
// all functions that start like "func(gr *GorillaRouter)" are like methods of the GorillaRouter Class and "gr" is the equivalent of "this" in Java
// the same lower/upper case rules apply to methods (private and public)
type GorillaRouter struct {
	router  *mux.Router
	logPath string
	srv     *http.Server
}

func NewGorillaRouter() *GorillaRouter {
	gr := new(GorillaRouter)
	gr.router = mux.NewRouter()
	return gr
}

func (gr *GorillaRouter) GetRouter() *mux.Router {
	return gr.router
}

func (gr *GorillaRouter) AddRoute(path, methods string, handler func(w http.ResponseWriter, r *http.Request)) {
	gr.router.HandleFunc(path, handler).Methods(methods)
}

func (gr *GorillaRouter) Serve(hostAndIpBinding string) error {
	gr.srv = &http.Server{
		Addr: hostAndIpBinding,
		// timeouts set to protect against Slowloris attacks
		WriteTimeout: time.Minute * 5,
		ReadTimeout:  time.Minute * 5,
		IdleTimeout:  time.Minute * 5,
		Handler:      gr.router,
	}
	// a go routine (the thing with go keyword) creates a "lightweight" thread to run a function in parallel
	// to serve to a client ip, we create a go routine (with the anonymous function)
	go func() {
		log.Println("INFO [router/gorrilarouter.go] Starting HTTP service:", hostAndIpBinding)
		err := gr.srv.ListenAndServe()
		if err != nil {
			log.Println("ERROR [router/gorrilarouter.go] Fail to start HTTP(S) service:", err.Error())
		}
	}()
	return nil
}

// Really good article on context: https://www.sohamkamani.com/golang/context-cancellation-and-values/
func (gr *GorillaRouter) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	gr.srv.Shutdown(ctx)
	return nil
}
