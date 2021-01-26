package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

// frequencyAnalysisCmd 引数で渡した文字列のアルファベットの頻度分析を行います
var frequencyAnalysisCmd = &cobra.Command{
	Use:   "fa",
	Short: "count the number of each characters",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		a := args[0]
		m := countChars(a)
		for _, k := range keySort(m) {
			fmt.Printf("%v: %v回\n", string(k), m[k])
		}
	},
}

// countChars 引数で与えた文字列にアルファベット（大文字小文字を区別）がそれぞれ何回出現するかをカウントします
func countChars(s string) map[string]int {
	result := make(map[string]int)

	for _, r := range s {
		result[string(r)] = result[string(r)] + 1
	}

	return result
}

// keySort 引数で与えられたmapのキーをソートします
func keySort(m map[string]int) []string {
	var keys []string
	for k := range m {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)

	return keys
}

func init() {
	CaesarCmd.AddCommand(frequencyAnalysisCmd)
}
