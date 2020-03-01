package main

import (
	"fmt"
)

func dynamicArray(N int32, queries [][]int32) []int32 {
	var lastAns int32 = 0
	returnArr := []int32{}
	seq := make([][]int32, N)

	for _, value := range queries {
		queryType := value[0]
		x := value[1]
		y := value[2]
		index := (x ^ lastAns) % N
		if queryType == 1 {
			seq[index] = append(seq[index], y)
		} else {
			var size = int32(len(seq[index]))
			lastAns = seq[index][y%size]
			returnArr = append(returnArr, lastAns)
		}
	}
	return returnArr
}

func main() {

	var a = [][]int32{{1, 0, 5}, {1, 1, 7}, {1, 0, 3}, {2, 1, 0}, {2, 1, 1}}
	ans := dynamicArray(3, a)
	fmt.Println(ans)
}
