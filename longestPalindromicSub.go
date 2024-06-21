package main

func longestPalindrome(s string) string {
	isPalindrome := func(sub string) bool {
		i := 0
		k := len(sub) - 1

		for i <= k {
			if sub[i] != sub[k] {
				return false
			}
			i++
			k--
		}
		return true
	}

	if isPalindrome(s) {
		return s
	}

	for window := len(s) - 1; window > 0; window-- {
		for i := 0; i+window <= len(s); i++ {
			wordWindow := s[i : i+window]
			if isPalindrome(wordWindow) {
				return wordWindow
			}
		}
	}

	return string(s[0])
}
