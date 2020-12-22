// Imagine we have a jar with numbers from one to ten in it.
// We take one number after the other out of the jar and put
// them in a line from left to right on the table.

package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
	"strconv"
)

const (
	runs = 10
	max  = runs
)

func getRandomInt(max int) (num int64) {
	num = 1
	// this is from 0 to max-1
	b, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	num += b.Int64()
	return num
}

func main() {
	kv := make(map[int]int64)

	i := 1
	for i <= runs {
		switch i {
		case 1:
			// shortcut for first key
			kv[i] = getRandomInt(max)
		case runs:
			// shortcut for last key
			var j int64
			for j = 1; j <= int64(runs); j++ {
				found := false
				for _, v := range kv {
					if j == v {
						found = true
						break
					}
				}
				if !found {
					kv[i] = j
					break
				}
			}
		default:
			for {
				num := getRandomInt(max)
				found := false
				for _, v := range kv {
					if num == v {
						found = true
						break
					}
				}
				if !found {
					kv[i] = num
					break
				}
			}
		}
		i++
	}

	var keys []int
	for k := range kv {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	var buf bytes.Buffer
	i = 1
	for _, k := range keys {
		buf.WriteString(strconv.FormatInt(kv[k], 10))
		if i < len(keys) {
			buf.WriteString(" ")
		}
		i++
	}

	fmt.Println(buf.String())
}
