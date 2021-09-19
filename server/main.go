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

type HandlersWrapper struct {
	db            interfaces.MongoInterface
	router        interfaces.Router
	hostIpBinding string
	frontEndPath  string
}

func NewApp() *HandlersWrapper {
	app := new(HandlersWrapper)

	app.router = router.NewGorillaRouter()
	app.hostIpBinding = os.Getenv("HOST_IP_BINDING")
	app.frontEndPath = os.Getenv("FRONT_END_PATH")

	login := new(database.MongoDBLogin)
	login.Uri = os.Getenv("DB_URI")
	app.db = database.NewMongoDB(login)

	return app
}

func (h *HandlersWrapper) Start() {
	h.db.Connect()

	fileServer := http.FileServer(http.Dir(h.frontEndPath))
	h.router.GetRouter().Path("/").Handler(fileServer)
	handlers.NewHandlers(h.router)

	h.router.Serve(h.hostIpBinding)
}

func (h *HandlersWrapper) Stop() {
	h.router.Stop()
	h.db.Stop()
}

func main() {
	app := NewApp()
	app.Start()
	defer app.Stop()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	os.Exit(0)
}
