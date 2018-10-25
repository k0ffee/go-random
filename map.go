// This creates a key-value storage, puts the numbers 1 to 10 as keys
// in there and adds a random value to each key. Then it copies the
// KV-store into a list of pairs, sorts according to the second column
// and puts the resulting first column into a string that is send to
// standard output.

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

const (
	runs = 10
	max  = runs * runs
)

var (
	i   int
	str string
)

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func getRandomInt(max int) (num int) {
	num = rand.Intn(max)
	return num
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	kv := make(map[int]int)

	i = 1
	for i <= runs {
		kv[i] = getRandomInt(max)
		i++
	}

	pairs := make(PairList, len(kv))

	i = 0
	for k, v := range kv {
		pairs[i] = Pair{k, v}
		i++
	}

	sort.Sort(pairs)

	str = ""
	for _, v := range pairs {
		str = fmt.Sprintf("%d ", v.Key) + str
	}

	fmt.Println(strings.TrimSpace(str))
}
