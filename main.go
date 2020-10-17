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

}
