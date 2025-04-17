package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		temp := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(temp[0])
		k, _ := strconv.Atoi(temp[1])

		scanner.Scan()
		inputArr := strings.Split(scanner.Text(), " ")
		inputIntArr := []int{}
		for _, val := range inputArr {
			newVal, _ := strconv.Atoi(val)
			inputIntArr = append(inputIntArr, newVal)
		}

		scanner.Scan()
		s := scanner.Text()
		fmt.Printf("%v, %v, %v, %v\n", n, k, s, inputIntArr)
	}
}

// compute power while applying modulo
// e.g. 2^5 with 998244353 mod would be mod_pow(2,5,998244353)
// func mod_pow(base int64, exp int, mod int64) int64 {
// 	var result int64 = 1
// 	base = base % mod
// 	for {
// 		if exp <= 0 {
// 			break
// 		}
// 		if exp%2 == 1 {
// 			result = (result * base) % mod
// 		}
// 		exp = exp >> 1
// 		base = (base * base) % mod
// 	}
// 	return result
// }
