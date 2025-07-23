package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var version = "1.0.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of DevLama",
	Long:  "Print the version number of DevLama CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("DevLama v%s\n", version)
	},
}

// getCurrentTime returns the current time formatted nicely
func getCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
