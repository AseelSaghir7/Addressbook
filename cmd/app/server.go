package app

import (
	"context"
	"github.com/addressBook/pkg/config"
	"github.com/addressBook/pkg/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func runApp(cmd *cobra.Command) {

	// parse config file
	const flag = "config"
	configFile, err := cmd.Flags().GetString(flag)
	if err != nil {
		log.Fatalf("unable to read command line flag, err : %v", err)
	}

	// loading config file
	c, err := config.Load(configFile)
	if err != nil {
		log.Fatalf("unable to read config file, err : %v", err)
	}

	s := server.New(c)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// running main server routine
	go s.Run()

	<-done
	log.Println("Stop signal received, now shutting down server ...")

	// handling server failure/stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	s.Stop(ctx)
}

func NewSPNCommand(in io.Reader, out, err io.Writer) *cobra.Command {

	cmds := &cobra.Command{
		Use:   "ab",
		Short: "ab: Address Book API server",
		Long:  "ab: Start Address Book API server with configuration file",
		Run: func(cmd *cobra.Command, args []string) {
			runApp(cmd)
		},
	}

	cmds.Flags().StringP("config", "c", "/etc/ab/config.yaml", "The path to the configuration file")

	return cmds
}
