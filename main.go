package main

import "fmt"

// AreAnagrams reports whether b is a rearrangement of a. One counting array:
// a's bytes increment, b's decrement, anagrams cancel to all zeros.
// shortcut: byte-level counts (ASCII/Latin-1); count runes for full Unicode.
func AreAnagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	var count [256]int
	for i := 0; i < len(a); i++ {
		count[a[i]]++
		count[b[i]]--
	}
	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(AreAnagrams("listen", "silent"), AreAnagrams("rat", "car"))
}
