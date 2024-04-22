package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`\B[?!,.]`)

func isGreater(k []string, m map[string]int, i, j int) bool {
	if m[k[i]] > m[k[j]] {
		return true
	} else if m[k[i]] == m[k[j]] {
		if k[i] < k[j] {
			return true
		}
	}
	return false
}

func Top10(s string) []string {
	const sliceLen int = 10
	counter := make(map[string]int)
	keys := make([]string, 0, len(counter))

	s = re.ReplaceAllString(s, "")

	for _, word := range strings.Fields(s) {
		w := strings.ToLower(word)
		if w != "" && w != "-" {
			counter[w]++
		}
	}

	for key := range counter {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return isGreater(keys, counter, i, j) })

	if len(keys) >= sliceLen {
		return keys[:sliceLen]
	}
	return keys
}
