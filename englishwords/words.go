package main

import (
	"fmt"
	"bufio"
	"net/http"
	"log"
)

func main(){
	// Get the list of words from a web resource using http.Get
	resp, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln("Error reading web resource:", err)
	}

	defer resp.Body.Close()  // resp.Body should be closed when done
	// Create a map of strings to integers
	wordCount := make(map[string]int)
	// We want to create a dictionary that counts how many words begin with each letter
	// The list of words contains all lowercase characters so we won't have to convert upper- to lowercase
	// In order to read the resp.Body we need a scanner from the bufio package
	sc := bufio.NewScanner(resp.Body)
	for sc.Scan() {
		word := sc.Text()
		firstLetter := string(word[0])
		wordCount[firstLetter]++
	}
	// Print out the wordcount in alphabetical order
	for i := 97; i <= 122; i++ {
		fmt.Printf("%v: %d\n", string(i), wordCount[string(i)])
	}
}
