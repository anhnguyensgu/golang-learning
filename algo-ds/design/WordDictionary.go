package main

import "fmt"

type WordDictionary struct {
	mapDict map[string]int
}

func Constructor() WordDictionary {
	initialMap := map[string]int{}
	return WordDictionary{mapDict: initialMap}
}

func (this *WordDictionary) AddWord(word string) {
	this.mapDict[word] = 1
}

func (this *WordDictionary) Search(word string) bool {
	_, result := this.mapDict[word]
	return result
}

func main() {
	wordDictionary := Constructor()

	wordDictionary.AddWord("abc")
	fmt.Println(wordDictionary.Search("abc"))
	fmt.Println(wordDictionary.Search("cde"))
}
