package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type SortByAge []user

func (t SortByAge) Len() int {
	return len(t)
}

func (t SortByAge) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t SortByAge) Less(i, j int) bool {
	return t[i].Age < t[j].Age
}

type SortByName []user

func (t SortByName) Len() int {
	return len(t)
}

func (t SortByName) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t SortByName) Less(i, j int) bool {
	return t[i].Last < t[j].Last
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	fmt.Println("Sorted by age-------------------------------")
	sbAge := SortByAge(users)
	sort.Sort(sbAge)
	for i, v := range sbAge {
		fmt.Println("Person nr:", i+1, " ", v.First, v.Last, v.Age)
		fmt.Println("Says:")
		for j, say := range v.Sayings {
			fmt.Println(j+1, ". ", say)
		}
	}

	fmt.Println("Sorted by name------------------------------")
	sbName := SortByName(users)
	sort.Sort(sbName)
	for i, v := range sbName {
		fmt.Println("Person nr:", i+1, " ", v.First, v.Last, v.Age)
		fmt.Println("Says:")
		for j, say := range v.Sayings {
			fmt.Println(j+1, ". ", say)
		}
	}
}
