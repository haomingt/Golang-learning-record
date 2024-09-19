package prose

import "strings"

func JoinWithCommas(phrase []string) string {
	var result string
	if len(phrase) == 2 {
		result = strings.Join(phrase, " and ")
	} else if len(phrase) == 3 {
		result = strings.Join(phrase[:len(phrase)-1], ",")
		result += ", and "
		result += phrase[len(phrase)-1]

	} else if len(phrase) == 1{
		result = phrase[0]
	}
	return result
}
