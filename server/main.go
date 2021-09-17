package main

import (
	"net/http"
	"os"
	"server/router"
)

const (
	ROUTER_LOG_PATH   = "/server/logs/router.txt"
	DATABASE_LOG_PATH = "/server/logs/database.txt"
)

type Engine struct {
	// TODO: add db
	router        router.Router
	hostIpBinding string
	frontEndPath  string
}

func NewEngine() *Engine {
	landingRepo := os.Getenv("LANDING_REPO")
	engine := new(Engine)
	engine.router = router.NewGorillaRouter(landingRepo + ROUTER_LOG_PATH)
	engine.hostIpBinding = os.Getenv("HOST_IP_BINDING")
	engine.frontEndPath = landingRepo + os.Getenv("FRONT_END_PATH")
	return engine
}

func (e *Engine) Start() {
	fileServer := http.FileServer(http.Dir(e.frontEndPath))
	e.router.GetRouter().PathPrefix("/").Handler(fileServer)
	e.router.Serve(e.hostIpBinding)
}

func (e *Engine) Stop() {
	// TODO: add db
	e.router.Stop()
}

func main() {
	engine := NewEngine()
	engine.Start()
	// TODO: implement a "wait for ctrl c" and then call defer engine.Stop()
}
