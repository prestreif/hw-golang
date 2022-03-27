package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type FreqStr struct {
	Word string
	Freq int
}

func getSliceStrings(fStr []FreqStr) []string { // get slice strings from slice FreqStr
	rStr := make([]string, 0, 10)
	for _, v := range fStr {
		rStr = append(rStr, v.Word)
	}
	return rStr
}

func Top10(inStr string) []string {
	aStr := strings.Fields(inStr)
	aFreq := make([]FreqStr, 0, len(aStr)/4)

	fAlpSort := func(i int, j int) bool { // function for first sort before count freq
		return aStr[i] < aStr[j]
	}
	sort.Slice(aStr, fAlpSort) // Мне кажется, что лишняя сортировка, будет лучше чем хэш

	for i := 0; i < len(aStr); i++ { // count freq
		if i > 0 && aStr[i-1] == aStr[i] {
			aFreq[len(aFreq)-1].Freq++
		} else {
			aFreq = append(aFreq, FreqStr{aStr[i], 1})
		}
	}

	fFreqSort := func(i int, j int) bool { // function for sort with freq and alph
		if aFreq[i].Freq == aFreq[j].Freq {
			return aFreq[i].Word < aFreq[j].Word
		}
		return aFreq[i].Freq > aFreq[j].Freq
	}
	sort.Slice(aFreq, fFreqSort)

	if len(aFreq) > 9 { // We need only TOP10
		aFreq = aFreq[:10]
	}
	return getSliceStrings(aFreq)
}
