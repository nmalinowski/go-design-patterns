package main

import "fmt"

// newPublication is a factory function that creates the specified publication type
func newPublication(pubType string, name string, pg int, pub string) (iPublication, error) {
	// DONE: Create the right kind of publication based on the given type
	if pubType == "newspaper" {
		return createNewspaper(name, pg, pub), nil
	}
	if pubType == "magazine" {
		return createMagazine(name, pg, pub), nil
	}
	// EXAMPLE END
	return nil, fmt.Errorf("No such publication type")
}
