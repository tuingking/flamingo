package cli

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/tuingking/flamingo/config"
	"github.com/tuingking/flamingo/infra/logger"
	"github.com/tuingking/flamingo/internal/product"
)

type CLI struct {
	Cmd *cobra.Command
	Cfg *config.Config
	Log *logger.Logger

	ProductSvc product.Service
}

func (c *CLI) Execute() {
	c.initCmd()
	c.initCsvCmd()
	c.initMigrateCmd()

	if err := c.Cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func (c *CLI) initCmd() {
	c.Cmd = &cobra.Command{
		Use:   "fla",
		Short: "flamingo CLI tools",
	}
}
