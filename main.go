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
	input := flag.String("input", "うどん食べ ステーキ食べて 寿司食べる", "string to be change the vertically")
	flag.Parse()

	words := strings.Fields(*input)
	reverse(words)

	// 縦書きに変換
	for i := 0; i < maxLength(words); i++ {
		line := ""
		for _, word := range words {
			runes := []rune(word)
			if i < len(runes) {
				if string(runes[i]) == "ー" {
					line += "｜"
				} else {
					line += string(runes[i])
				}

			} else {
				line += "　" // 文字がない場合は全角スペースで埋める
			}
			line += " " // 行は半角スペースで区切る
		}
		fmt.Println(line)
	}
	return nil
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
