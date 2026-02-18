package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

var (
	input string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "goiku",
		Short: "HAIKU maker - 俳句を縦書き風に表示します",
		Long: `goiku は俳句を縦書き風に表示するツールです。
スペースで区切られた単語を縦書き風に変換して出力します。

例:
  goiku --input "うどん食べ ステーキ食べて 寿司食べる"`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}

	// TODO: I want auto-generation mode.
	rootCmd.Flags().StringVarP(&input, "input", "i", "うどん食べ ステーキ食べて 寿司食べる", "縦書き風に変換する文字列")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	words := strings.Fields(input)
	reverse(words)
	convertToVertical(words)

	return nil
}

func convertToVertical(words []string) {
	for i := 0; i < maxLength(words); i++ {
		line := buildVerticalLine(words, i)
		fmt.Println(line)
	}
}

func buildVerticalLine(words []string, lineIndex int) string {
	var line strings.Builder
	for _, word := range words {
		runes := []rune(word)
		if lineIndex < len(runes) {
			line.WriteString(convertProlongedSoundMark(runes[lineIndex]))
		} else {
			line.WriteString("　") // 文字がない場合は全角スペースで埋める
		}
		line.WriteString(" ") // 行は半角スペースで区切る
	}
	return line.String()
}

func convertProlongedSoundMark(r rune) string {
	if r == 'ー' {
		return "｜"
	}
	return string(r)
}

func maxLength(words []string) int {
	max := 0
	for _, s := range words {
		i := utf8.RuneCountInString(s)
		if i > max {
			max = i
		}
	}
	return max
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
