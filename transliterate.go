package thaana

import (
	"bytes"
	"strings"
)

// https://github.com/jawish/thaana_conversions_php/blob/49b17bbf70c7f025899149d8bd912c1f8a66b5aa/ThaanaConversions.php#L49
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
	')': 0041, '(': 0040, 'Q': 65010,
}

// https://github.com/jawish/thaana_conversions_php/blob/49b17bbf70c7f025899149d8bd912c1f8a66b5aa/ThaanaConversions.php#L103
func AsciiToUnicodeNumbers(s string) []int {
	var unicodes []int = []int{}
	for _, v := range bytes.Split([]byte(s), []byte("")) {
		if i, ok := mapAsciiToUnicode[v[0]]; ok {
			unicodes = append(unicodes, i)
		} else {
			unicodes = append(unicodes, int(v[0]))
		}
	}
	return unicodes
}

// https://github.com/jawish/thaana_conversions_php/blob/49b17bbf70c7f025899149d8bd912c1f8a66b5aa/ThaanaConversions.php#L156
func UnicodeNumbersToUtf(s []int) string {
	var unicodes = strings.Builder{}
	length := len(s) - 1
	for i := range s {
		unicodes.WriteString(string(rune(s[length-i])))
	}
	return unicodes.String()
}

// https://github.com/jawish/thaana_conversions_php/blob/49b17bbf70c7f025899149d8bd912c1f8a66b5aa/ThaanaConversions.php#L156
func AsciiToUnicode(s string) string {
	return UnicodeNumbersToUtf(AsciiToUnicodeNumbers(s))
}
