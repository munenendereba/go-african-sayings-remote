package main

import "fmt"

func main() {
	fmt.Println("Would you like to see specific sayings or all? Press 'Enter' to see for all languages: ")
	fmt.Print("Languages available: ")

	availableLangs := GetAvailableLanguages()

	fmt.Println(availableLangs)

	var language string
	fmt.Scanln(&language)

	foundLang := false

	for _, b := range availableLangs.Languages {
		if b.Lang == language {
			foundLang = true
		}
	}

	if !foundLang {
		fmt.Printf("Language %v not found, we will fetch all quotes.", language)
		language = ""
	}

	fmt.Println("")

	fmt.Println("See one quote or all quotes? Enter 1 for one quote and press 'Enter' to see all quotes.")
	var showAllSayings string
	fmt.Scanln(&showAllSayings)

	var showAll bool = true

	if showAllSayings == "1" {
		showAll = false
	}

	AfricanSaying(language, showAll)

	fmt.Println("Press 'Enter' to exit: ")
	fmt.Scanln()
}
