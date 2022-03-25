package main

import (
	"fmt"
	"sort"
)

/*
Go’s sort package implements sorting for builtins and user-defined types
*/
func main() {
	strs := []string{"d", "c", "a", "b"}
	/*Sort methods are specific to the builtin type;
	here’s an example for strings
	*/
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	ints := []int{10, 398, 4, 87}
	sort.Ints(ints)
	fmt.Println("Ints", ints)
	/*
		We can also use sort to check if a slice is already in sorted order
	*/
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
