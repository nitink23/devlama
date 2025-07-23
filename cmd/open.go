package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open [file|directory|url]",
	Short: "Open files, directories, or URLs with default system applications",
	Long: `Open command allows you to open:
- Files with their default applications
- Directories in file explorer
- URLs in the default web browser

Examples:
  devlama open file.txt          # Opens file.txt with default text editor
  devlama open .                 # Opens current directory in file explorer
  devlama open /path/to/folder   # Opens specific folder
  devlama open https://google.com # Opens URL in default browser`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		target := args[0]

		if err := openTarget(target); err != nil {
			fmt.Fprintf(os.Stderr, "Error opening %s: %v\n", target, err)
			os.Exit(1)
		}

		fmt.Printf("Opened: %s\n", target)
	},
}

func openTarget(target string) error {
	// Check if it's a URL
	if strings.HasPrefix(target, "http://") || strings.HasPrefix(target, "https://") {
		return openURL(target)
	}

	// Check if it's a file or directory
	if _, err := os.Stat(target); err == nil {
		return openFileOrDir(target)
	}

	// If it doesn't exist as a file/dir, treat it as a URL (maybe without protocol)
	if !strings.Contains(target, "://") {
		target = "https://" + target
	}

	return openURL(target)
}

func openFileOrDir(path string) error {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", absPath)
	case "darwin":
		cmd = exec.Command("open", absPath)
	case "linux":
		cmd = exec.Command("xdg-open", absPath)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	return cmd.Start()
}

func openURL(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	return cmd.Start()
}

func init() {
	rootCmd.AddCommand(openCmd)
}
