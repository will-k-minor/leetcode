package main

import "bytes"

func letterCombinations(digits string) []string {
	// have a map of each T9 value
	// maximum combinations of X arrays (each amonunting to <= 4)
	// can reduce the combinations b/c the order matters
	/**
	  maximum length is len(string)

	  recursion

	  each recursive loop keeps track of the current combo and what is left of the digits

	  if we have no more digits left, push the current combo to results
	*/

	var results []string

	t9 := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	// need to use a buffer to append to a string
	var backtrack func(currentCombo bytes.Buffer, remainingDigits string)

	backtrack = func(currentCombo bytes.Buffer, remainingDigits string) {
		if len(remainingDigits) == 0 {
			if currentCombo.Len() > 0 {
				results = append(results, currentCombo.String())
			}
			return
		}

		letters, exists := t9[remainingDigits[0]]

		if exists {
			for _, letter := range letters {
				currentCombo.Write([]byte(letter))
				// take the current combo, and run it against a reduced digits array
				backtrack(currentCombo, remainingDigits[1:])
				// in order for currentCombo not to rack up over each loop, we need to truncate it
				currentCombo.Truncate(currentCombo.Len() - 1)
			}
		}

		return

	}

	backtrack(bytes.Buffer{}, digits)
	return results
}
