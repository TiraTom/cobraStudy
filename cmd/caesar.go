package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CaesarCmd シーザー暗号を解読するコマンド
var CaesarCmd = &cobra.Command{
	Use:   "Caesar",
	Short: "Break the Caesar cipher",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Caesar")
	},
}
