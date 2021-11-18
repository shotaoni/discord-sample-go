package main

import (
	// "fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"sample/x/mux"
)

// Session is declared in the global space so it can be easily used
// throughout this program.
// In this use case, there is no error that would be returned.
var Session, _ = discordgo.New()

// Read in all configuration options from both environment variables and
// command line arguments.
var Router = mux.New()

func init() {
	log.Printf("hit!")

	// Discord Authentication Token
	Session.Token = "Bot TOKEN"
	Session.AddHandler(Router.OnMessageCreate)

	// register route
	Router.Route("help", "Display this message.", Router.Help)
	Router.Route("joke", "send a joke.", Router.Joke)
}

func main() {
	// Open a websocket connection to Discord
	err := Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}

	// Wait for a CTRL-C
	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Clean up
	Session.Close()
}
