package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// ASCIICODEa　ASCII_CODEのaの値
// ASCIICODEz　ASCII_CODEのzの値
const (
	ASCIICODEa = 97
	ASCIICODEz = 122
)

// CaesarCmd シーザー暗号を解読するコマンド
var CaesarCmd = &cobra.Command{
	Use:   "caesar",
	Short: "Break the Caesar cipher",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		lowerCaseRunes := []rune(strings.ToLower(args[0]))

		for i := 1; i < 26; i++ {
			var convertedASCIICodes []int
			for _, r := range lowerCaseRunes {
				var convertedASCIICode int

				n := int(r)

				switch {
				case n-i < ASCIICODEa:
					// TODO このロジックミスった
					convertedASCIICode = ASCIICODEz - i - (n - ASCIICODEa) + 1
				default:
					convertedASCIICode = n - i
				}

				convertedASCIICodes = append(convertedASCIICodes, convertedASCIICode)
			}

			var convertedMessage string = ""
			for _, r := range convertedASCIICodes {
				convertedMessage = convertedMessage + string(rune(r))
			}

			fmt.Println(convertedMessage)
		}

		return nil
	},
}
