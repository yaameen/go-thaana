package thaana

import (
	"strings"
)

// Complete ASCII to Unicode mapping for Thaana script
var mapAsciiToUnicode = map[byte]int{
	'h': 1920, 'S': 1921, 'n': 1922, 'r': 1923,
	'b': 1924, 'L': 1925, 'k': 1926, 'a': 1927,
	'v': 1928, 'm': 1929, 'f': 1930, 'd': 1931,
	't': 1932, 'l': 1933, 'g': 1934, 'N': 1935,
	's': 1936, 'D': 1937, 'z': 1938, 'T': 1939,
	'y': 1940, 'p': 1941, 'j': 1942, 'C': 1943,
	'X': 1944, 'H': 1945, 'K': 1946, 'J': 1947,
	'R': 1948, 'x': 1949, 'B': 1950, 'F': 1951,
	'Y': 1952, 'Z': 1953, 'A': 1954, 'G': 1955,
	'q': 1956, 'V': 1957, 'w': 1958, 'W': 1959,
	'i': 1960, 'I': 1961, 'u': 1962, 'U': 1963,
	'e': 1964, 'E': 1965, 'o': 1966, 'O': 1967,
	'c': 1968, ',': 1548, ';': 1563, '?': 1567,
	'(': 40, ')': 41, 'M': 1929, 'P': 1941, 'Q': 65010,
}

// IsASCIIOnly checks if string contains only ASCII characters
func IsASCIIOnly(s string) bool {
	for _, b := range []byte(s) {
		if b > 127 {
			return false
		}
	}
	return true
}

// HasTransliterationChars checks if ASCII string has transliteration characters
func HasTransliterationChars(s string) bool {
	for _, b := range []byte(s) {
		if _, ok := mapAsciiToUnicode[b]; ok {
			return true
		}
	}
	return false
}

// ConvertASCIIWord converts a single ASCII word to Thaana
func ConvertASCIIWord(word string) string {
	if !IsASCIIOnly(word) {
		return word
	}

	if !HasTransliterationChars(word) {
		return word
	}

	// Convert to unicode numbers
	var unicodes []int
	for _, char := range []byte(word) {
		if mapped, ok := mapAsciiToUnicode[char]; ok {
			unicodes = append(unicodes, mapped)
		} else {
			unicodes = append(unicodes, int(char))
		}
	}

	// Reverse and convert to string (RTL)
	var chars []rune
	length := len(unicodes) - 1
	for i := range unicodes {
		chars = append(chars, rune(unicodes[length-i]))
	}

	// Fix parentheses for RTL
	for i, char := range chars {
		if char == '(' {
			chars[i] = ')'
		} else if char == ')' {
			chars[i] = '('
		}
	}

	return string(chars)
}

// AsciiToUnicodeNumbers converts ASCII string to array of Unicode code points
func AsciiToUnicodeNumbers(s string) []int {
	var unicodes []int
	for _, char := range []byte(s) {
		if mapped, ok := mapAsciiToUnicode[char]; ok {
			unicodes = append(unicodes, mapped)
		} else {
			unicodes = append(unicodes, int(char))
		}
	}
	return unicodes
}

// UnicodeNumbersToUtf converts array of Unicode code points to UTF-8 string
func UnicodeNumbersToUtf(s []int) string {
	var chars []rune
	length := len(s) - 1

	for i := range s {
		chars = append(chars, rune(s[length-i]))
	}

	// Fix parentheses for RTL
	for i, char := range chars {
		if char == '(' {
			chars[i] = ')'
		} else if char == ')' {
			chars[i] = '('
		}
	}

	return string(chars)
}

// AsciiToUnicode converts ASCII transliteration to Thaana Unicode.
// Handles mixed content: only ASCII segments are transliterated; existing
// Unicode (e.g. Thaana) is left unchanged.
func AsciiToUnicode(s string) string {
	if s == "" {
		return s
	}
	if IsASCIIOnly(s) {
		return UnicodeNumbersToUtf(AsciiToUnicodeNumbers(s))
	}
	var result strings.Builder
	var asciiBuf strings.Builder
	for _, r := range s {
		if r < 128 {
			asciiBuf.WriteRune(r)
		} else {
			if asciiBuf.Len() > 0 {
				converted := UnicodeNumbersToUtf(AsciiToUnicodeNumbers(asciiBuf.String()))
				asciiBuf.Reset()
				result.WriteString(strings.TrimLeft(converted, " "))
				result.WriteRune(' ')
			}
			result.WriteRune(r)
		}
	}
	if asciiBuf.Len() > 0 {
		result.WriteString(UnicodeNumbersToUtf(AsciiToUnicodeNumbers(asciiBuf.String())))
	}
	return result.String()
}

// SafeAsciiToUnicode processes mixed content word by word
func SafeAsciiToUnicode(s string) string {
	if len(s) == 0 {
		return s
	}

	words := strings.Fields(s)
	var result []string

	for _, word := range words {
		if IsASCIIOnly(word) && HasTransliterationChars(word) {
			converted := ConvertASCIIWord(word)
			result = append(result, converted)
		} else {
			result = append(result, word)
		}
	}

	return strings.Join(result, " ")
}
