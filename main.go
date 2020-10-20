package main

import (
	"bufio"
	"fmt"
	"os"
)

func printDiff(diff []MarkdownDiff) {
	fmt.Printf("\033[1m%d action(s) staged\033[0m\n", len(diff))
	for _, i := range diff {
		switch i.Action {
		case Create:
			fmt.Printf("\033[32mCREATE\033[0m\t%s\n", i.File.Path)
		case Update:
			fmt.Printf("\033[33mUPDATE\033[0m\t%s\n", i.File.Path)
		case Delete:
			fmt.Printf("\033[31mDELETE\033[0m\t%s\n", i.File.Path)
		}
	}
}

func main() {
	localFiles, err := LocalArticleState(".")
	if err != nil {
		panic(err.Error())
	}
	remoteFiles, err := RemoteArticleState("alice", "")
	if err != nil {
		panic(err.Error())
	}
	diff := diffMarkdownFiles(localFiles, remoteFiles)
	printDiff(diff)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
