package cmd

import (
	"fmt"
	"runtime"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devlama",
	Short: "DevLama - A developer productivity CLI tool",
	Long: `DevLama is a CLI application designed to boost developer productivity.
Built with Go and Cobra, it provides various utilities and tools
for developers in their daily workflow.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Show banner when no command is specified
		showBanner()
	},
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}

func showBanner() {
	// ASCII Art for DevLama
	banner := `
██████╗ ███████╗██╗   ██╗██╗      █████╗ ███╗   ███╗ █████╗ 
██╔══██╗██╔════╝██║   ██║██║     ██╔══██╗████╗ ████║██╔══██╗
██║  ██║█████╗  ██║   ██║██║     ███████║██╔████╔██║███████║
██║  ██║██╔══╝  ╚██╗ ██╔╝██║     ██╔══██║██║╚██╔╝██║██╔══██║
██████╔╝███████╗ ╚████╔╝ ███████╗██║  ██║██║ ╚═╝ ██║██║  ██║
╚═════╝ ╚══════╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝
`

	// Colors (ANSI escape codes)
	cyan := "\033[36m"
	yellow := "\033[33m"
	green := "\033[32m"
	blue := "\033[34m"
	magenta := "\033[35m"
	reset := "\033[0m"
	bold := "\033[1m"

	// Print banner with colors
	fmt.Print(cyan + banner + reset)

	// Tool info
	fmt.Printf("\n%s%s🚀 Developer Productivity CLI Tool%s\n", bold, yellow, reset)
	fmt.Printf("%s%sVersion:%s %s%s%s\n", bold, blue, reset, green, version, reset)
	fmt.Printf("%s%sGo Version:%s %s%s%s\n", bold, blue, reset, green, runtime.Version(), reset)
	fmt.Printf("%s%sPlatform:%s %s%s/%s%s\n", bold, blue, reset, green, runtime.GOOS, runtime.GOARCH, reset)
	fmt.Printf("%s%sTime:%s %s%s%s\n", bold, blue, reset, green, time.Now().Format("2006-01-02 15:04:05"), reset)

	// Quick help
	fmt.Printf("\n%s%s📋 Quick Commands:%s\n", bold, magenta, reset)
	fmt.Printf("  %s• devlama hello --name <name>%s   - Greet someone\n", cyan, reset)
	fmt.Printf("  %s• devlama open <file|url>%s       - Open files/URLs\n", cyan, reset)
	fmt.Printf("  %s• devlama version%s               - Show version\n", cyan, reset)
	fmt.Printf("  %s• devlama --help%s                - Show all commands\n", cyan, reset)

	// Footer
	fmt.Printf("\n%s%s💡 Ready to boost your productivity!%s\n\n", bold, yellow, reset)
}

func init() {
	// Global flags can be added here
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose output")
}
