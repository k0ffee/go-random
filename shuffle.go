// I've just found this in the "math/rand" package.

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	numbers := strings.Fields("1 2 3 4 5 6 7 8 9 10")
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	fmt.Println(strings.Join(numbers, " "))
}
