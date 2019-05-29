// 数字を受け取ってでかくする
package main

import "os"
import "fmt"

var one = [5][3]int {
	{0, 1, 0,},
	{0, 1, 0,},
	{0, 1, 0,},
	{0, 1, 0,},
	{0, 1, 0,},
}
var two = [5][3]int {
	{1, 1, 1,},
	{0, 0, 1,},
	{0, 1, 1,},
	{1, 0, 0,},
	{1, 1, 1,},
}

var three = [5][3]int {
	{1, 1, 1,},
	{0, 0, 1,},
	{0, 1, 0,},
	{0, 0, 1,},
	{1, 1, 1,},
}

var four = [5][3]int {
	{1, 0, 0,},
	{1, 0, 1,},
	{1, 1, 1,},
	{0, 0, 1,},
	{0, 0, 1,},
}

var five = [5][3]int {
	{1, 1, 1,},
	{1, 0, 0,},
	{1, 1, 1,},
	{0, 0, 1,},
	{1, 1, 1,},
}

var six = [5][3]int {
	{1, 1, 0,},
	{1, 0, 0,},
	{1, 1, 1,},
	{1, 0, 1,},
	{1, 1, 1,},
}

var seven = [5][3]int {
	{1, 1, 1,},
	{0, 0, 1,},
	{0, 0, 1,},
	{0, 1, 0,},
	{0, 1, 0,},
}

var eight = [5][3]int {
	{1, 1, 1,},
	{1, 0, 1,},
	{1, 1, 1,},
	{1, 0, 1,},
	{1, 1, 1,},
}

var nine = [5][3]int {
	{1, 1, 1,},
	{1, 0, 1,},
	{1, 1, 1,},
	{0, 0, 1,},
	{0, 1, 1,},
}

var zero = [5][3]int {
	{1, 1, 1,},
	{1, 0, 1,},
	{1, 0, 1,},
	{1, 0, 1,},
	{1, 1, 1,},
}

var digits = [10][5][3]int{ zero, one, two, three, four, five, six, seven, eight, nine }

func printRow(char [5][3]int, row int) {
	for col := 0; col <= 2; col++ {
		if (char[row][col] == 1) {
			fmt.Print("█")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Print(" ")
}

func main() {
	numbers := os.Args[1]
	for row := 0; row <= 4; row++ {
		for _, c := range numbers {
			switch c {
			case '0': printRow(digits[0], row)
			case '1': printRow(digits[1], row)
			case '2': printRow(digits[2], row)
			case '3': printRow(digits[3], row)
			case '4': printRow(digits[4], row)
			case '5': printRow(digits[5], row)
			case '6': printRow(digits[6], row)
			case '7': printRow(digits[7], row)
			case '8': printRow(digits[8], row)
			case '9': printRow(digits[9], row)
			default: fmt.Printf("otherwise [%c]\n", c)
			}
		}
		fmt.Println()
	}
}
