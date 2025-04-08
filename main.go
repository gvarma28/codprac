package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < t; i++ {
		scanner.Scan()
		temp := strings.Split(scanner.Text(), " ")
		n, _ := strconv.Atoi(temp[0])
		k, _ := strconv.Atoi(temp[1])

		scanner.Scan()
		s := scanner.Text()
		fmt.Printf("%v, %v, %v\n", n, k, s)
	}
}
