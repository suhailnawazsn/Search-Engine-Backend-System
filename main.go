package main

import (
	"fmt"
	"strings"
)

type Document struct {
	ID      int
	Content string
}

type InvertedIndex map[string][]int

func NewInvertedIndex() InvertedIndex {
	return make(InvertedIndex)
}

func (index InvertedIndex) Add(doc Document) {
	words := strings.Fields(strings.ToLower(doc.Content))
	for _, word := range words {
		index[word] = append(index[word], doc.ID)
	}
}

func (index InvertedIndex) Search(query string) []int {
	words := strings.Fields(strings.ToLower(query))
	if len(words) == 0 {
		return nil
	}

	results := make(map[int]int)
	for _, word := range words {
		if docIDs, found := index[word]; found {
			for _, id := range docIDs {
				results[id]++
			}
		}
	}

	var finalResults []int
	for id, count := range results {
		if count == len(words) {
			finalResults = append(finalResults, id)
		}
	}
	return finalResults
}

func main() {
	docs := []Document{
		{ID: 1, Content: "Go is an open-source programming language."},
		{ID: 2, Content: "Golang is great for concurrent programming."},
		{ID: 3, Content: "Go provides a rich standard library."},
	}

	index := NewInvertedIndex()
	for _, doc := range docs {
		index.Add(doc)
	}

	// Search for documents
	query := "Go programming"
	results := index.Search(query)

	fmt.Printf("Documents matching query '%s': %v\n", query, results)
}
