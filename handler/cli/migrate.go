package cli

import (
	"github.com/spf13/cobra"
	"github.com/tuingking/flamingo/infra/mysql"
	"github.com/tuingking/flamingo/internal/jade"
)

var (
	migrateCmd *cobra.Command
)

// update the struct here...
var obj = jade.Page{}

func (c *CLI) initMigrateCmd() {
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "migrate tools",
		Run: func(cmd *cobra.Command, args []string) {
			mysql.Migrate(obj)
		},
	}

	c.Cmd.AddCommand(migrateCmd)
	migrateCmd.Flags().IntP("number", "n", 0, "generate csv file with n product")
}
