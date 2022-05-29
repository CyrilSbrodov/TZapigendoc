package internal

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"TZapigendoc/data"
)

type Replace struct {
	Use      string `json:"use"`
	Text     string `json:"text"`
	Recordid int    `json:"recordid"`
}

type Result struct {
	URLPDF  string `json:"urlpdf"`
	URLWord string `json:"urlword"`
}

func ReplaceWord(replace *Replace) (string, string) {
	file := data.FileName()
	stringNeeded := "    <w:t>_</w:t>"
	word := "_"
	stringToReplace := replace.Text

	f, err := os.Open(file.FilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		text := scanner.Text()
		if scanner.Text() == stringNeeded {
			words := strings.Split(scanner.Text(), "")
			result := ""
			for i := 0; i < len(words); i++ {
				if words[i] == word {
					words[i] = stringToReplace
				}
				result += words[i]
			}
			text = result
		}
		lines = append(lines, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(file.FilePathNew, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	URLWord := "http://localhost:8282/" + file.FilePathNew
	fmt.Println(URLWord)
	URLPDF := "no adr"
	return URLWord, URLPDF
}
