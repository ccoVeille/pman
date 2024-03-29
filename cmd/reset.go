package cmd

import (
	"fmt"
	"pman/pkg"

	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Deletes the current indexed projects , run pman init to reindex the projects",
	Run: func(cmd *cobra.Command, args []string) {
		err := pkg.DeleteDb()
		if err != nil {
			fmt.Println(err)
		}
        fmt.Println("Successfully reset the database , run pman init to reindex the projects")
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}