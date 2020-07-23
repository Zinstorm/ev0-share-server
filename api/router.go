package api

import (
	"encoding/json"
	"ev0/bot"
	"fmt"
	"net/http"
	"os"

	h "github.com/gorilla/handlers"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ShareGameServer struct {
	Username   string `json:"username"`
	APIKey     string `json:"apiKey"`
	ServerIP   string `json:"serverIp"`
	ServerName string `json:"serverName"`
	ServerMap  string `json:"serverMap"`
}

var decoder = schema.NewDecoder()

//Route defines a route to with the handler and if it should use a authenication
type Route struct {
	Handler http.HandlerFunc
}

//App defines the database and all the handlers of the served application
type App struct {
	dg       *bot.DiscordSession
	handlers map[string]Route
}

//NewApp is used in the server.go to start the api
func NewApp(dg *bot.DiscordSession, cors bool) App {
	app := App{
		dg:       dg,
		handlers: make(map[string]Route),
	}
	shareServer := app.ShareServer

	if !cors {
		shareServer = disableCors(shareServer)
	}
	app.handlers["/api/share"] = Route{Handler: shareServer}

	return app
}

func (a *App) ShareServer(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
	var shareGameServer ShareGameServer
	err = decoder.Decode(&shareGameServer, r.PostForm)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
	//log.Printf("Server Info: %+v", shareGameServer)
	if shareGameServer.APIKey == viper.GetString("apikey") {
		a.dg.Session.ChannelMessageSend(viper.GetString("shareServer.channelId"), fmt.Sprintf("Join %s on %s: steam://connect/%s \n %s", shareGameServer.Username, shareGameServer.ServerMap, shareGameServer.ServerIP, shareGameServer.ServerName))
		w.Write([]byte("OK"))
		return
	}
	w.Write([]byte("FUCK YOU"))
}

func (a *App) Serve() error {
	for path, route := range a.handlers {
		http.Handle(path, (route.Handler))
	}

	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", h.CompressHandler(h.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}
