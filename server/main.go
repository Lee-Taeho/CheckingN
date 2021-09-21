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

type App struct {
	db            interfaces.MongoInterface
	router        interfaces.Router
	hostIpBinding string
	frontEndPath  string
}

func NewApp() *App {
	app := new(App)

	app.router = router.NewGorillaRouter()
	app.hostIpBinding = os.Getenv("HOST_IP_BINDING")
	app.frontEndPath = os.Getenv("FRONT_END_PATH")

	login := new(database.MongoDBLogin)
	login.Uri = os.Getenv("DB_URI")
	app.db = database.NewMongoDB(login)

	return app
}

func (a *App) Start() {
	a.db.Connect()

	fileServer := http.FileServer(http.Dir(a.frontEndPath))
	a.router.GetRouter().Path("/").Handler(fileServer)
	a.router.GetRouter().PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(a.frontEndPath+"/assets/"))))

	handlers.NewHandlers(a.router)

	a.router.Serve(a.hostIpBinding)
}

func (a *App) Stop() {
	a.router.Stop()
	a.db.Stop()
}

func main() {
	app := NewApp()
	app.Start()
	defer app.Stop()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
