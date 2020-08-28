package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build project as a LaTeX document",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
