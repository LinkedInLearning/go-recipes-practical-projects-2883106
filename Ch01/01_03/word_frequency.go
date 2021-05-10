package main

import (
	"fmt"
)

func main() {
	fmt.Println(frequency(moby))
}

func frequency(words []string) map[string]int {
	freq := make(map[string]int)
	for _, w := range words {
		freq[w]++
	}

	return freq
}

var moby = []string{
	"call", "me", "ishmael", "some", "years", "ago", "never", "mind", "how",
	"long", "precisely", "having", "little", "or", "no", "money", "in", "my",
	"purse", "and", "nothing", "particular", "to", "interest", "me", "on",
	"shore", "i", "thought", "i", "would", "sail", "about", "a", "little",
	"and", "see", "the", "watery", "part", "of", "the", "world", "it", "is",
	"a", "way", "i", "have", "of", "driving", "off", "the", "spleen", "and",
	"regulating", "the", "circulation", "whenever", "i", "find", "myself",
	"growing", "grim", "about", "the", "mouth", "whenever", "it", "is", "a",
	"damp", "drizzly", "november", "in", "my", "soul", "whenever", "i", "find",
	"myself", "involuntarily", "pausing", "before", "coffin", "warehouses",
	"and", "bringing", "up", "the", "rear", "of", "every", "funeral", "i",
	"meet", "and", "especially", "whenever", "my", "hypos", "get", "such",
	"an", "upper", "hand", "of", "me", "that", "it", "requires", "a", "strong",
	"moral", "principle", "to", "prevent", "me", "from", "deliberately",
	"stepping", "into", "the", "street", "and", "methodically", "knocking",
	"people", "s", "hats", "off", "then", "i", "account", "it", "high", "time",
	"to", "get", "to", "sea", "as", "soon", "as", "i", "can", "this", "is",
	"my", "substitute", "for", "pistol", "and", "ball", "with", "a",
	"philosophical", "flourish", "cato", "throws", "himself", "upon", "his",
	"sword", "i", "quietly", "take", "to", "the", "ship", "there", "is",
	"nothing", "surprising", "in", "this", "if", "they", "but", "knew", "it",
	"almost", "all", "men", "in", "their", "degree", "some", "time", "or",
	"other", "cherish", "very", "nearly", "the", "same", "feelings", "towards",
	"the", "ocean", "with", "me",
}
