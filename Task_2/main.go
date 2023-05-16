package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 9119
	str := ""
	digit := 0
	for {
		digit = num % 10
		str += strconv.Itoa(digit * digit)
		num = num / 10
		if num == 0 {
			break
		}
	}

	numNew, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	fmt.Println(numNew)
}
