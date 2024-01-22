/*
Copyright © 2024 Kevin Park <krapi0314@gmail.com>
*/

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/krapie/showbox/server"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs showbox API server for show ticketing.",
	Long:  `Runs showbox API server for show ticketing.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := server.Run()
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
