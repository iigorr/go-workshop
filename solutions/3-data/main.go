package main

import (
	"./addressbook"
	"./importer"
	"fmt"
)

func main() {
	ab := addressbook.New()

	ab.Add(addressbook.NewEntry("Gopher", "Google Street"))
	ab.Add(addressbook.NewEntry("Igor", "igor@lankin.de"))
	ab.Add(addressbook.NewEntry("Lord Voldemort", "lordy64@hotmail.com"))

	ab.Delete("Lord Voldemort")

	fmt.Println("Printing address book:")
	for _, entry := range ab.List() {
		fmt.Printf("%v: %v\n", entry.Name(), entry.Address())
	}
	fmt.Println()

	//variant 2
	imp := importer.New() // "generates" entries

	//variant 3
	//imp := importer.NewGithubImporter()  // not implemented

	imp.Import(ab) //imports some entries into ab

	fmt.Println("Persons with 'go' in name:")
	for _, entry := range ab.Find("go") {
		fmt.Printf("%v: %v\n", entry.Name(), entry.Address())
	}
}
