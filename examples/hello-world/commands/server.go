package commands

import (
	"fmt"
	stdlog "log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/josegonzalez/cli-skeleton/command"
	"github.com/josegonzalez/service-skeleton/server"
	"github.com/posener/complete"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"
)

func init() {
	log.Logger = server.DefaultLogger()
	stdlog.SetFlags(0)
	stdlog.SetOutput(log.Logger)
}

type ServerCommand struct {
	command.Meta
}

func (c *ServerCommand) Name() string {
	return "server"
}

func (c *ServerCommand) Synopsis() string {
	return "Server command"
}

func (c *ServerCommand) Help() string {
	return command.CommandHelp(c)
}

func (c *ServerCommand) Examples() map[string]string {
	appName := os.Getenv("CLI_APP_NAME")
	return map[string]string{
		"Starts a server": fmt.Sprintf("%s %s", appName, c.Name()),
	}
}

func (c *ServerCommand) Arguments() []command.Argument {
	args := []command.Argument{}
	return args
}

func (c *ServerCommand) AutocompleteArgs() complete.Predictor {
	return complete.PredictNothing
}

func (c *ServerCommand) ParsedArguments(args []string) (map[string]command.Argument, error) {
	return command.ParseArguments(args, c.Arguments())
}

func (c *ServerCommand) FlagSet() *flag.FlagSet {
	f := c.Meta.FlagSet(c.Name(), command.FlagSetClient)
	return f
}

func (c *ServerCommand) AutocompleteFlags() complete.Flags {
	return command.MergeAutocompleteFlags(
		c.Meta.AutocompleteFlags(command.FlagSetClient),
		complete.Flags{},
	)
}

func (c *ServerCommand) Run(args []string) int {
	flags := c.FlagSet()
	flags.Usage = func() { c.Ui.Output(c.Help()) }
	if err := flags.Parse(args); err != nil {
		c.Ui.Error(err.Error())
		c.Ui.Error(command.CommandErrorText(c))
		return 1
	}

	_, err := c.ParsedArguments(flags.Args())
	if err != nil {
		c.Ui.Error(err.Error())
		c.Ui.Error(command.CommandErrorText(c))
		return 1
	}

	if err := runServer(); err != nil {
		c.Ui.Error(err.Error())
		return 1
	}

	return 0
}

func runServer() error {
	server.InitializeLogging()
	r := server.GetServer()
	// add your custom logic here
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
	})

	return server.StartServer(r)
}
