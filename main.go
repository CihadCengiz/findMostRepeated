package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// WordCount holds word and count pair
type WordCount struct {
	word  string
	count int
}

func Print(text string, w io.Writer) {
	fmt.Fprintln(w, text)
}

func main() {
	//Get input from command line
	fmt.Print("Enter text: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	//Split the input string into an array
	words := strings.Split(input, ",")

	//Count same words in words
	m := make(map[string]int)
	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	//Create and fill slice of word-count pairs for sorting by count
	wordCounts := make([]WordCount, 0, len(m))
	for key, val := range m {
		wordCounts = append(wordCounts, WordCount{word: key, count: val})
	}

	//Sort wordCount slice by decreasing count number
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].count > wordCounts[j].count
	})

	//Display the most repeated word
	Print(wordCounts[0].word, os.Stdout)
}
