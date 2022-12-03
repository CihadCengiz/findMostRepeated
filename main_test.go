package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	// input := strings.NewReader("[\"change\",\"random\",\"\",\"\",\"\",\"\"]") //returns "" -> bug
	// input := strings.NewReader("[\"change\",\"pie\",\"apple\",\"red\",\"red\",\"red\"]")
	input := strings.NewReader("[0,-1,-5,-5,null,null,null,null]")
	// input := strings.NewReader("[\"\",\"\",\"\",\"\",\"\",\"\"]")
	// input := strings.NewReader("[\"red\",\"red\",\"red\",\"red\",\"red\",\"red\"]")
	// input := strings.NewReader("[\"apple\",\"apple\",\"apple\",\"most\",\"most\",\"most\"]") //returns apple -> bug

	scanner := bufio.NewScanner(input)
	msg := "Input Array:"
	fmt.Fprintln(os.Stdout, msg)

	scanner.Scan()
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}

	//Get scanner text and split it into words
	words := scanner.Text()
	wordsArray := strings.Split(words, ",")

	//Count same words in words
	m := make(map[string]int)
	for _, word := range wordsArray {
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

	//Check if the input is empty
	if len(wordsArray) == 0 {
		t.Log("empty input")
	}
	//Print the most repeated word
	t.Logf("Most repeated: %s\n", wordCounts[0].word)
}
