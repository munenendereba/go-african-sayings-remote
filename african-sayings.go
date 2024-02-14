package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

type Languages struct {
	Languages []Lang `json:"languages"`
}

type Lang struct {
	Lang string `json:"lang"`
}

type Saying struct {
	Saying      string `json:"saying"`
	Translation string `json:"translation"`
}

type Sayings struct {
	Sayings []Saying `json:"sayings"`
}

var baseUrl string = "https://raw.githubusercontent.com/munenendereba/african-sayings/main/"

func GetAvailableLanguages() Languages {
	resp, err := http.Get(baseUrl + "languages.json")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	lang, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var languages Languages

	json.Unmarshal(lang, &languages)

	return languages
}

func AfricanSaying(langIn string, allSayings bool) {

	fileSuffix := "-sayings.json"

	if langIn != "" {
		langFilePath := baseUrl + langIn + fileSuffix

		sayings, err := GetSayings(langFilePath)
		if err != nil {
			log.Fatalln(err)
		}

		if allSayings {
			fmt.Println(sayings)
		} else {
			sayingPos := rand.Intn(len(sayings.Sayings))

			fmt.Println(sayings.Sayings[sayingPos].Saying)
		}

	} else {
		languages := GetAvailableLanguages()

		if !allSayings {
			randLang := len(languages.Languages)
			langIn = languages.Languages[rand.Intn(randLang)].Lang

			sayingsUrl := baseUrl + langIn + fileSuffix

			sayings, err := GetSayings(sayingsUrl)

			if err != nil {
				log.Fatalln(err)
			}

			sayingPos := rand.Intn(len(sayings.Sayings))

			sayingDisplay := fmt.Sprintf("%+v", sayings.Sayings[sayingPos])

			fmt.Println(sayingDisplay)

		} else {

			for i := 0; i < len(languages.Languages); i++ {
				sayingsUrl := baseUrl + languages.Languages[i].Lang + fileSuffix

				sayings, err := GetSayings(sayingsUrl)

				if err != nil {
					log.Fatalln(err)
				}

				sayingsDisplay := fmt.Sprintf("%+v", sayings)

				fmt.Println(sayingsDisplay)
			}
		}
	}
}

func GetSayings(sayingsUrl string) (Sayings, error) {
	resp, err := http.Get(sayingsUrl)

	if err != nil {
		log.Fatalln(err.Error())
		return Sayings{}, err
	}

	fmt.Println("")
	fmt.Println("Fetched", sayingsUrl, "successfully.")

	defer resp.Body.Close()

	fileBytesValue, err2 := io.ReadAll(resp.Body)

	if err2 != nil {
		log.Fatalln(err2.Error())

		return Sayings{}, err2
	}

	var sayings Sayings

	errUnmarshall := json.Unmarshal(fileBytesValue, &sayings)

	if errUnmarshall != nil {
		log.Fatalln(errUnmarshall.Error())

		return Sayings{}, errUnmarshall
	}

	return sayings, nil
}
