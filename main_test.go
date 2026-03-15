package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "three elements",
			input: []string{"a", "b", "c"},
			want:  []string{"c", "b", "a"},
		},
		{
			name:  "single element",
			input: []string{"a"},
			want:  []string{"a"},
		},
		{
			name:  "empty",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "haiku words",
			input: []string{"うどん食べ", "ステーキ食べて", "寿司食べる"},
			want:  []string{"寿司食べる", "ステーキ食べて", "うどん食べ"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.input)
			if len(tt.input) != len(tt.want) {
				t.Fatalf("reverse() len = %d, want %d", len(tt.input), len(tt.want))
			}
			for i, v := range tt.input {
				if v != tt.want[i] {
					t.Errorf("reverse()[%d] = %q, want %q", i, v, tt.want[i])
				}
			}
		})
	}
}

func TestMaxLength(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  int
	}{
		{
			name:  "haiku words",
			input: []string{"寿司食べる", "ステーキ食べて", "うどん食べ"},
			want:  7,
		},
		{
			name:  "empty",
			input: []string{},
			want:  0,
		},
		{
			name:  "single word",
			input: []string{"あいう"},
			want:  3,
		},
		{
			name:  "same length",
			input: []string{"ab", "cd"},
			want:  2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxLength(tt.input)
			if got != tt.want {
				t.Errorf("maxLength() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestConvertProlongedSoundMark(t *testing.T) {
	tests := []struct {
		name  string
		input rune
		want  string
	}{
		{
			name:  "prolonged sound mark",
			input: 'ー',
			want:  "｜",
		},
		{
			name:  "regular hiragana",
			input: 'あ',
			want:  "あ",
		},
		{
			name:  "regular katakana",
			input: 'ス',
			want:  "ス",
		},
		{
			name:  "ascii character",
			input: 'a',
			want:  "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertProlongedSoundMark(tt.input)
			if got != tt.want {
				t.Errorf("convertProlongedSoundMark(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestBuildVerticalLine(t *testing.T) {
	words := []string{"寿司食べる", "ステーキ食べて", "うどん食べ"}
	tests := []struct {
		name      string
		lineIndex int
		want      string
	}{
		{
			name:      "first line",
			lineIndex: 0,
			want:      "寿 ス う ",
		},
		{
			name:      "second line",
			lineIndex: 1,
			want:      "司 テ ど ",
		},
		{
			name:      "line with prolonged sound mark",
			lineIndex: 2,
			want:      "食 ｜ ん ",
		},
		{
			name:      "line with padding",
			lineIndex: 5,
			want:      "　 べ 　 ",
		},
		{
			name:      "last line",
			lineIndex: 6,
			want:      "　 て 　 ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildVerticalLine(words, tt.lineIndex)
			if got != tt.want {
				t.Errorf("buildVerticalLine(%v, %d) = %q, want %q", words, tt.lineIndex, got, tt.want)
			}
		})
	}
}
