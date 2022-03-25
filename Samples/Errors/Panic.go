package main

import "os"

func main() {
	/*
		use panic throughout this site to check for unexpected errors
	*/
	panic("A Problem Occured:The system cannot find the path specified")
	/*A common use of panic is to abort if a function returns an error
	 */
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
	/*
		When first panic in main fires, the program exits without reaching the rest of the code
	*/

}
