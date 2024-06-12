package main

import (
	"fmt"
	"sort"
	"strings"
)

type custom struct {
	field         string
	accessedValue bool
	id            int
}

type custArray []custom

func (c custArray) Len() int { return len(c) }
func (c custArray) Swap(i, j int) {
	if c[i].accessedValue && c[j].accessedValue {
		c[i], c[j] = c[j], c[i]
	}
}
func (c custArray) Less(i, j int) bool {
	return strings.Compare(c[i].field, c[j].field) == -1
}

func (c custArray) String() string {
	return fmt.Sprintf("erdiocane ")
}

func main() {
	var el interface{}

	el = "123"

	s, ok := el.(int)
	if ok {
		print(s)
	} else {
		print("zio ma che cazzo fai, mi hai passato un %T", el)
	}
	elemets := []custom{{"a", false, 0}, {"ff", true, 1}, {"gbd", true, 2}, {"cc", true, 3}, {"dm", false, 4}}
	fmt.Println(elemets)
	sort.Sort(custArray(elemets))
	fmt.Println(elemets)
	print(custArray(elemets))
}
