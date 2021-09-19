package main

import (
	"net/http"
	"os"
	"os/signal"
	"server/database"
	"server/handlers"
	"server/middleware/interfaces"
	"server/router"
	"syscall"
)

type EngineWrapper struct {
	db            interfaces.MongoInterface
	router        interfaces.Router
	hostIpBinding string
	frontEndPath  string
}

func NewEngine() *EngineWrapper {
	engine := new(EngineWrapper)

	engine.router = router.NewGorillaRouter()
	engine.hostIpBinding = os.Getenv("HOST_IP_BINDING")
	engine.frontEndPath = os.Getenv("FRONT_END_PATH")

	mongoInfo := new(database.MongoInfo)
	mongoInfo.Uri = os.Getenv("DB_URI")
	engine.db = database.NewMongo(mongoInfo)

	return engine
}

func (e *EngineWrapper) Start() {
	e.db.Connect()

	fileServer := http.FileServer(http.Dir(e.frontEndPath))
	e.router.GetRouter().Path("/").Handler(fileServer)
	handlers.NewEngine(e.router)

	e.router.Serve(e.hostIpBinding)
}

func (e *EngineWrapper) Stop() {
	e.router.Stop()
	e.db.Stop()
}

func main() {
	engine := NewEngine()
	engine.Start()
	defer engine.Stop()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	os.Exit(0)
}
