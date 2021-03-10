package numbertowords

import (
	"errors"
)

//maxNumber is the maximum allowed number to convert
const MaxNumber = 99999

//Convert converts numbers to words
func Convert(number int) (string, error) {
	if number < 0 || number > MaxNumber {
		return "", errors.New("error occurred")
	}

	words := [21]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen",
	}

	tenwords := [10]string{
		"",
		"",
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety",
	}

	result := ""

	thousands := number / 1000

	if thousands > 0 {
		tens := thousands / 10
		units := thousands % 10
		if tens < 2 {
			result = words[thousands] + " thousand "
		} else if units == 0 {
			result = tenwords[tens] + " thousand "
		} else {
			result = tenwords[tens] + " " + words[units] + " thousand "
		}
		number = number % 1000
		if number == 0 {
			return result, nil
		}

	}

	hundreds := number / 100

	if hundreds > 0 {
		result = result + words[hundreds] + " hundred "
		number = number % 100
		if number == 0 {
			return result, nil
		}
	}

	tens := number / 10
	units := number % 10

	if tens < 2 {
		return result + words[number], nil
	}

	if units == 0 {
		return result + tenwords[tens], nil
	}

	return result + tenwords[tens] + " " + words[units], nil

}
