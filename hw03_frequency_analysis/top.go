package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const punctuation = `,.'"!?`

func Top10(input string) []string {
	if input == "" {
		return []string{}
	}

	tokens := strings.Fields(input)

	freqTable := CreateFreqTable(tokens)

	words := SortByFreqAndLexical(freqTable)

	if len(words) < 10 {
		return words
	}
	return words[:10]
}

func SortByFreqAndLexical(freqTable map[string]int) []string {
	words := make([]string, 0, len(freqTable))
	for key := range freqTable {
		words = append(words, key)
	}

	sort.SliceStable(words, func(i int, j int) bool {
		return words[i] < words[j]
	})
	sort.SliceStable(words, func(i int, j int) bool {
		return freqTable[words[i]] > freqTable[words[j]] // требуется убывание кол-ва
	})

	return words
}

func CreateFreqTable(tokens []string) map[string]int {
	freqTable := make(map[string]int, len(tokens)/2)
	for _, token := range tokens {
		if token != "-" {
			freqTable[Normalize(token)]++
		}
	}
	return freqTable
}

func Normalize(token string) string {
	var result string
	result = strings.Trim(token, punctuation)
	result = strings.ToLower(result)
	return result
}
