package main

import (
	"fmt"
	"sort"
)

// type TablicaI []int

// func (t TablicaI) Len() int {
// 	return len(t)
// }

// func (t TablicaI) Swap(i, j int) {
// 	t[i], t[j] = t[j], t[i]
// }

// func (t TablicaI) Less(i, j int) bool {
// 	return t[i] < t[j]
// }

// type TablicaS []string

// func (t TablicaS) Len() int {
// 	return len(t)
// }

// func (t TablicaS) Swap(i, j int) {
// 	t[i], t[j] = t[j], t[i]
// }

// func (t TablicaS) Less(i, j int) bool {
// 	return t[i] < t[j]
// }

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	// tab := TablicaI(xi)
	// sort.Sort(tab)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	// tabs := TablicaS(xs)
	// sort.Sort(tabs)
	sort.Strings(xs)
	fmt.Println(xs)
}
