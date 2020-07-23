package main

import (
	"ev0/api"
	"ev0/bot"
	"ev0/config"
	"ev0/logging"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func init() {
	logging.Init()
	config.Init()
}

func main() {
	dg, err := bot.Init()
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}

	app := api.NewApp(dg, false)
	err = app.Serve()

	// Wait here until CTRL-C or other term signal is received.
	log.Info("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	// Cleanly close down the Discord session.
	dg.Session.Close()
}
