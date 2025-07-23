package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var name string

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello to someone",
	Long:  "A simple command that greets a person by name",
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			fmt.Println("Hello, World!")
		} else {
			fmt.Printf("Hello, %s!\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().StringVarP(&name, "name", "n", "", "Name to greet")
} 