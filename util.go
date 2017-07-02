package kit

import "strings"

func CamelCase(input string) string {
	if len(input) != 0 {
		input = strings.ToLower((input)[0:1]) + (input)[1:]
	}
	return input
}

func PascalCase(input string) string {
	if len(input) != 0 {
		input = strings.ToUpper((input)[0:1]) + (input)[1:]
	}
	return input
}

func LowerMapKey(data map[string][]string) {
	for k, v := range data {
		delete(data, k)
		data[strings.ToLower(k)] = v
	}
}
