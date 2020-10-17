package main

import (
	"fmt"
)

func sliceContains(a MarkdownFile, b []MarkdownFile) bool {
	for _, i := range b {
		if a.Path == i.Path {
			return true
		}
	}
	return false
}

func diffMarkdownFiles(a []MarkdownFile, b []MarkdownFile) (diff []MarkdownFile) {
	// Loop two times, first to find ∉b, second to find ∉a
	for i := 0; i < 2; i++ {
		for _, j := range a {
			found := sliceContains(j, b)
			if !found {
				diff = append(diff, j)
			}
		}
		if i == 0 {
			a, b = b, a // Swap the slices
		}
	}
	return
}

func main() {
	localFiles, err := LocalArticleState(".")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%d local markdown files\n", len(localFiles))
	fmt.Println(localFiles)
	remoteFiles, err := RemoteArticleState("alice", "")
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%d remote markdown files\n", len(remoteFiles))
	fmt.Println(remoteFiles)
	diff := diffMarkdownFiles(remoteFiles, localFiles)
	fmt.Printf("%d files different\n", len(diff))
	fmt.Println(diff)
}
