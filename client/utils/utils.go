package utils

// Instr returns true if the character char is in the string str. Otherwise false
func Instr(str, char string) bool {
	for _, c := range str {
		if char == string(c) {
			return true
		}
	}
	return false
}
