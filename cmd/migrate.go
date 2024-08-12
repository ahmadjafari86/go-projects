package cmd

import (
	"blog/pkg/bootstrap"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Tables migrate to database",
	Long:  "Tables migrate to database",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func migrate() {
	bootstrap.Migrate()
}
