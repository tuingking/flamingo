package cli

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	csvCmd *cobra.Command
)

func (c *CLI) initCsvCmd() {
	csvCmd = &cobra.Command{
		Use:   "csv",
		Short: "csv tools",
		Run: func(cmd *cobra.Command, args []string) {
			number, err := cmd.Flags().GetInt("number")
			if err != nil {
				logrus.Error(err)
				return
			}

			_, err = c.ProductSvc.GenerateProductsCsv(cmd.Context(), number)
			if err != nil {
				logrus.Error(err)
				return
			}
		},
	}

	c.Cmd.AddCommand(csvCmd)
	csvCmd.Flags().IntP("number", "n", 0, "generate csv file with n product")
}
