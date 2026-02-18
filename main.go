package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	if err := Main(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func Main() error {
	// TODO: I want auto-generation mode.
	input := flag.String("input", "うどん食べ ステーキ食べて 寿司食べる", "string to be change the vertically")
	flag.Parse()

	words := strings.Fields(*input)
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
