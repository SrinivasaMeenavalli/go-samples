package main

/*
Go offers built-in support for regular expressions
*/

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	/*
		Tests whether a pattern matches a string
	*/
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println("Matching:", match)
	//Compile an optimized Regexp struct.
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println("Match String:", r.MatchString("peach"))
	//finds the match for the regexp
	fmt.Println("Find String:", r.FindString(" blah blah peach punch"))
	/*
		finds the first match but returns the start and end indexes
	*/
	fmt.Println("idx:", r.FindStringIndex("peach punch"))
	/*
		The Submatch variants include information
		about both the whole-pattern matches
		and the submatches within those matches.
		For example this will return
		information for both p([a-z]+)ch and ([a-z]+)
	*/
	fmt.Println(r.FindStringSubmatch("peach punch"))
	/*The All variants of these functions
	apply to all matches in the input, not just the first
	*/

	fmt.Println(r.FindAllString("peach punch pinch", -1))
	//return information about the indexes of matches and submatches
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))
	/*Providing a non-negative integer
	as the second argument to these functions
	 will limit the number of matches
	*/
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	//[]byte argument
	fmt.Println(r.Match([]byte("peach")))
	/*MustCompile panics instead of returning
	an error, which makes it safer to use for global variables*/

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)
	/*The regexp package can also be used
	to replace subsets of strings with other values*/

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
