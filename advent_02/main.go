package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	arr := stringConvert()

	for p0 := 0; p0 <= len(arr); p0 += 4  {
		opcode := arr[p0]
		if opcode == 99 {
			break
		}

		outputLocation := arr[p0+3]
		input1 := arr[arr[p0+1]]
		input2 := arr[arr[p0+2]]
		value := 0

		switch opcode {
		case 1:
			value = input1 + input2
		case 2:
			value = input1 * input2
		default:

		}

		arr[outputLocation] = value
	}

	fmt.Println(arr)
}

func stringConvert() []int {
	a := strings.Split(input, ",")
	b := make([]int, len(input))
	for i, v := range a {
		b[i], _ = strconv.Atoi(v)
	}

	return b
}

var input = `1,12,2,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,6,19,1,19,5,23,2,13,23,27,1,10,27,31,2,6,31,35,1,9,35,39,2,10,39,43,1,43,9,47,1,47,9,51,2,10,51,55,1,55,9,59,1,59,5,63,1,63,6,67,2,6,67,71,2,10,71,75,1,75,5,79,1,9,79,83,2,83,10,87,1,87,6,91,1,13,91,95,2,10,95,99,1,99,6,103,2,13,103,107,1,107,2,111,1,111,9,0,99,2,14,0,0`

//var input = `1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,1,6,19,1,19,5,23,2,13,23,27,1,10,27,31,2,6,31,35,1,9,35,39,2,10,39,43,1,43,9,47,1,47,9,51,2,10,51,55,1,55,9,59,1,59,5,63,1,63,6,67,2,6,67,71,2,10,71,75,1,75,5,79,1,9,79,83,2,83,10,87,1,87,6,91,1,13,91,95,2,10,95,99,1,99,6,103,2,13,103,107,1,107,2,111,1,111,9,0,99,2,14,0,0`
