package cmd

import (
	"fmt"

	"github.com/drpaneas/vimm-cli/internal/snes"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Retrieve new list of roms",
	Long:  `Retrieve new list of roms`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
		snes.RunIt()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
