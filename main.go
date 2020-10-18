package main

import (
	"fmt"
)

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
	diff := diffMarkdownFiles(localFiles, remoteFiles)
	fmt.Printf("%d files different\n", len(diff))
	fmt.Println(diff)
}
