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

// 2 given m, find if n exists, 1^3...(n-1)^3+n^3
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

func FindNb(m int) int{
	for n:=1,m>0,n++{
		m-= n*n*n
		if m==0{
			return n
		}
	}
	return -1
}


// 3: given 3 sides, find if it can make a triangle

func IsTriangle(a, b, c int) bool {
  if a+b >c && b+c > a && c+a>b {
    return true
    }else{
    return false
    }
}


// 4: mumbling

func Accum(s string) string {
    // your code
   	var str string
	for i:=0;i<len(s);i++{
	str+=strings.ToUpper(string(s[i]))+strings.Repeat(strings.ToLower(string(s[i])),i)
	if(i+1!=len(s)){
		str+=string("-")
	}
	}
    return str
}