package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "usage",
		Short: "简易用法",
		Long:  "便捷查询的简易用法",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln("err")
	}
}

func init() {
	rootCmd.Version = Version
}
