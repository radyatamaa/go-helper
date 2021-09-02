package array

import "strings"

func InArray(str string, list []string) bool {
	str = strings.ToLower(str)
	for _, v := range list {
		if strings.ToLower(v) == str {
			return true
		}
	}
	return false
}

func GetLastofCharacter(countChar int,char string) string {
	char = char[len(char)-countChar:]
	return char
}
