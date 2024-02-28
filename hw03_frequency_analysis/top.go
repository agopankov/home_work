package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	if len(text) == 0 {
		return nil
	}

	words := strings.Fields(text)

	freqMap := make(map[string]int)
	for _, word := range words {
		freqMap[word]++
	}

	type wordFreq struct {
		word  string
		count int
	}

	freqWords := make([]wordFreq, 0, len(freqMap))
	for word, count := range freqMap {
		freqWords = append(freqWords, wordFreq{word, count})
	}

	sort.Slice(freqWords, func(i, j int) bool {
		if freqWords[i].count == freqWords[j].count {
			return freqWords[i].word < freqWords[j].word
		}
		return freqWords[i].count > freqWords[j].count
	})

	var top10 []string
	for i := 0; i < len(freqWords) && i < 10; i++ {
		top10 = append(top10, freqWords[i].word)
	}

	return top10
}
