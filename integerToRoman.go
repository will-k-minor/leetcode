package main

func IntegerToRoman(num int) string {
	// Run a modulus on each Roman against num
	// keep going until we reach 0

	romanNumerals := [][]interface{}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	remainder := num
	result := ""
	for _, numeral := range romanNumerals {
		numeralValue := numeral[0].(int)
		for remainder >= numeralValue {
			result += numeral[1].(string)
			remainder -= numeralValue
		}
	}

	return result
}
