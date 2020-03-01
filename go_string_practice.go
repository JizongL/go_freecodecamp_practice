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

// Remove first and last charactere
func RemoveChar(word string) string {
	return word[1:len(word)-1]
  }

// Counting Duplicates, given words, find number of distinct duplicated letters

import("strings")
func duplicate_count(s1 string) int {
    var total int
	  records:=map[string]int{}
	  for i:=0;i<len(s1);i++{		
		  records[strings.ToLower(string(s1[i]))] +=1		  
		  if records[strings.ToLower(string(s1[i]))]==2{
			  total+=1
		    }		
		  }
		return total
	}
	
// better solution than last one, but use the same principle.  
func duplicate_count(s string) (c int) {
    h := map[rune]int{}
    for _, r := range strings.ToLower(s) {
      if h[r]++; h[r] == 2 { c++ }
    }
    return
}

// Given two strings, combine them and removed duplicate and return sorted
// my first approach, turned it's too complicated
import(

	"fmt"
	"strings"
	"sort"
	)

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }
func StringToRuneSlice(s string) []rune {
      var r []rune
      for _, runeValue := range s {
              r = append(r, runeValue)
      }
      return r
}


func TwoToOne(s1 string, s2 string) (str string) {
    h := map[rune]int{}
    for _, r := range strings.ToLower(s1) {
      if h[r]++; h[r] == 1 { str+=string(r) }
    }
    for _, r := range strings.ToLower(s2) {
      if h[r]++; h[r] == 1 { str+=string(r) }
    }

	var r ByRune = StringToRuneSlice(str)
	sort.Sort(r)				
    return string(r)
}


func main() {
	a := "xyaabbbccccdefww"
	b := "xxxxyyyyabklmopq"
	//c:="abcdefghijklmnopqrstuvwxyz"
	d:= duplicate_count(a,b)
	fmt.Println(d)
}


// better solution 
func TwoToOne(s1 string, s2 string)string{
	chars := strings.Split(s1+s2,"")
	sort.Strings(chars)
	result:= ""
	for _,s:= range chars{
		chr:= string(s)
		if !strings.Contains(result,chr){
			result+=chr
		}
	}
	return result
}