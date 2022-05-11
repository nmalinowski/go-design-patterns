package main

import "fmt"

// Iterate using a callback function
func main() {
	// DONE: use IterateBooks to iterate via a callback function
	lib.IterateBooks(myBookCallback)

	// DONE: Use IterateBooks to iterate via anonymous function
	lib.IterateBooks(func(b Book) error {
		fmt.Println("Book Author:", b.Author)
		return nil
	})
	// DONE: create a BookIterator
	iter := lib.createIterator()
	for iter.hasNext() {
		book := iter.next()
		fmt.Printf("Book %+v\n", book)
	}
}

// END EXAMPLE

// This callback function processes an individual Book object
func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
