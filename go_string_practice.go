// 02/28/2020
// codewar practice go fundemental

// 1: DNA to RNA translate
package kata

import (
	"strings"
)

func DNAtoRNA(dna string) string {
	// your code here
	return strings.Replace(dna, "T", "U", -1)
}

func DNAtoRNA2(dna string) string {
	var res string
	for _, letter := range dna {
		if letter == 'T' {
			letter = 'U'
		}
		res += string(letter)
	}
	return res
}

// 2
func FindNb(m int) int {

	n := float64(0)
	var m2 float64 = 0
	for m2 < float64(m) {
		n += 1
		m2 += math.Pow(n, float64(3))

	}
	if int(m2) == m {
		return int(n)
	} else {
		return -1
	}
}
