package main

import "testing"

func TestAreAnagrams(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"listen", "silent", true},
		{"rat", "car", false},
		{"", "", true},
		{"a", "ab", false},
		{"aab", "abb", false}, // same letters, different multiplicity
		{"anagram", "nagaram", true},
	}
	for _, c := range cases {
		if got := AreAnagrams(c.a, c.b); got != c.want {
			t.Errorf("AreAnagrams(%q,%q) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
