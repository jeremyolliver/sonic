package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/devfacet/gocmd"
)

func main() {
	flags := struct {
		Help    bool `short:"h" long:"help" description:"Display usage" global:"true"`
		Version bool `short:"v" long:"version" description:"Display version"`
		Debug   struct {
			Settings bool `settings:"true" allow-unknown-arg:"true"`
		} `command:"debug" description:"Print arguments"`
		Info struct {
			Full     bool `short:"f" long:"full" required:"false" description:"show full output"`
			Settings bool `settings:"true" allow-unknown-arg:"true"`
		} `command:"info" description:"Show information about a compute resource"`
	}{}

	// Debug command
	gocmd.HandleFlag("Debug", func(cmd *gocmd.Cmd, args []string) error {
		DebugCommand()

		return nil
	})

	gocmd.HandleFlag("Info", func(cmd *gocmd.Cmd, args []string) error {
		InfoCommand(args[1], flags.Info.Full)

		return nil
	})

	// Init the app
	gocmd.New(gocmd.Options{
		Name:        "sonic",
		Version:     "1.0.0",
		Description: "Find and connect wherever you need to go",
		Flags:       &flags,
		ConfigType:  gocmd.ConfigTypeAuto,
	})
}

var sharedsession *session.Session = nil

// Enabling SharedConfigEnable, ensures defaults are loaded from ~/.aws/config as well as ~/.aws/credentials
// This ensures auto-discovery of items such as region, and profiles
func SharedAWSSession() *session.Session {
	if sharedsession != nil {
		return sharedsession
	} else {
		sharedsession = session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
	}
	return sharedsession
}
