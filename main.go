package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	author = "alice"
	token  = "token_a"
)

func printDiff(diff []AssetDiff) {
	fmt.Printf("\033[1m%d action(s) staged\033[0m\n", len(diff))
	for _, i := range diff {
		switch i.Action {
		case Create:
			fmt.Printf("\033[32mCREATE\033[0m\t%s\n", i.Asset.Path)
		case Update:
			fmt.Printf("\033[33mUPDATE\033[0m\t%s\n", i.Asset.Path)
		case Delete:
			fmt.Printf("\033[31mDELETE\033[0m\t%s\n", i.Asset.Path)
		}
	}
}

func main() {
	if err := LoadIgnore("."); err != nil {
		panic(err.Error())
	}
	localFiles, err := LocalAssetState(".")
	if err != nil {
		panic(err.Error())
	}
	remoteFiles, err := RemoteAssetState(author, token)
	if err != nil {
		panic(err.Error())
	}
	diff := diffAssets(localFiles, remoteFiles)
	printDiff(diff)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	sync(author, token, diff)
}
