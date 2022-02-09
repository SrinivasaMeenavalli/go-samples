package main

import "fmt"

/**
Maps are Go’s built-in associative data type
*/
func main() {
	//make(map[key-type]val-type)
	m := make(map[string]int)
	//Set key/value pairs using typical name[key] = val
	m["Key1"] = 6
	m["Key2"] = 8
	fmt.Println("Map", m)
	//Get a value for a key with name[key]
	value1 := m["Key1"]
	fmt.Println("Value 1:", value1)
	fmt.Println("Map length:", len(m))
	//delete removes key/value pairs from a map
	delete(m, "Key1")
	fmt.Println("Updated Map", m)
	/**The optional second return value when getting a
	value from a map indicates if the key was present in the map
	*/

	_, keyNotFound := m["Key1"]
	fmt.Println("prs:", keyNotFound)

	_, keyExisting := m["Key2"]
	fmt.Println("prs:", keyExisting)
}
