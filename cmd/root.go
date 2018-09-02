package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	stdout = log.New(os.Stdout, "", 0)
	stderr = log.New(os.Stderr, "shake", 0)
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "shake [OPTIONS] [TARGETS]",
	Short: "Automated target-based building and deployment system",
	Long: `Automated target-based building and deployment system`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
