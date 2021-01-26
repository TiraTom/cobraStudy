package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// caesar2Cmd <https://blog.amedama.jp/entry/2015/10/13/212155>参考にロジックを修正
var caesar2Cmd = &cobra.Command{
	Use:   "caesar2",
	Short: "better logic",
	Run: func(cmd *cobra.Command, args []string) {
		candidates := caesar2(args[0])
		for _, c := range candidates {
			fmt.Println(c)
		}
	},
}

func caesar2(arg string) []string {
	var results []string

	for i := 1; i < 26; i++ {
		var result string
		for _, r := range arg {
			result = result + string(shiftString(r, i))

		}
		results = append(results, result)
	}

	return results
}

func shiftString(r rune, shift int) rune {
	var result int
	b := int(r)

	switch {
	case ('A' <= b && b <= 'Z'):
		result = (b-'A'+shift)%26 + 'A'
	case ('a' <= b && b <= 'z'):
		result = (b-'a'+shift)%26 + 'a'
	default:
		result = b
	}
	return rune(result)
}

func init() {
	CaesarCmd.AddCommand(caesar2Cmd)
}
