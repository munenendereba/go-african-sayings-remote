package main

import "fmt"

func main() {
	fmt.Println("Would you like to see specific sayings or all? Enter to see for all languages: ")
	fmt.Print("Languages available: ")
	fmt.Println(GetAvailableLanguages())

	var language string
	fmt.Scanln(&language)

	fmt.Println("See one quote or all quotes? Enter 1 for one quote and press 'Enter' to see all quotes.")
	var showAllSayings string
	fmt.Scanln(&showAllSayings)

	var showAll bool = true

	if showAllSayings == "1" {
		showAll = false
	}

	AfricanSaying(language, showAll)
}
