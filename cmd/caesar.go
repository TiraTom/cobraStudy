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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(caesar(cmd, args))
	},
}

func caesar(cmd *cobra.Command, args []string) []string {
	lowerCaseRunes := []rune(strings.ToLower(args[0]))

	var result []string

	for i := 1; i < 26; i++ {
		result = append(result, shiftASCIICODEs(lowerCaseRunes, i))
	}

	return result
}

func shiftASCIICODEs(runes []rune, shift int) string {
	var convertedASCIICodes []int
	for _, r := range runes {
		convertedASCIICodes = append(convertedASCIICodes, shiftASCIICODE(r, shift))
	}

	var convertedString string = ""
	for _, r := range convertedASCIICodes {
		convertedString = convertedString + string(rune(r))
	}

	return convertedString
}

func shiftASCIICODE(base rune, shift int) int {
	var convertedASCIICode int

	n := int(base)

	switch {
	case n-shift < ASCIICODEa:
		convertedASCIICode = ASCIICODEz - ASCIICODEa - shift + n + 1
	default:
		convertedASCIICode = n - shift
	}

	return convertedASCIICode
}
