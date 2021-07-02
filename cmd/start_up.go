package cmd

import (
	"easy-search/config"
	"easy-search/db"
	"easy-search/models"
	"easy-search/router"
	"github.com/spf13/cobra"
)

func init() {
	runServer := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			config.LoadConfig()
			db.Init()
			err := db.Conn.AutoMigrate(
				&models.SenQq{},
			)
			if err != nil {
				panic(err)
			}
			router.Init()
		},
	}
	rootCmd.AddCommand(runServer)
}
