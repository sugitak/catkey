package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "catkey",
	Short: "catkey reads ssl/tls keys",
	Long:  "catkey reads ssl/tls keys",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("It works!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
