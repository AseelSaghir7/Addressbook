package main

import (
	"github.com/addressBook/cmd/app"
	log "github.com/sirupsen/logrus"
	"os"
)

func init(){
	log.SetOutput(os.Stdout)
}

func main() {

	command := app.NewSPNCommand(os.Stdin, os.Stdout, os.Stderr)

	// start server
	if err := command.Execute(); err != nil {
		log.Errorf("unable to start server, err : %v", err)
		os.Exit(1)
	}
}
