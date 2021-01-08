package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/moyen-blog/sync-dir/client"
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

func fatalError(message string, err error) {
	fmt.Println("\033[31mERROR\033[0m", message+":", err.Error())
	os.Exit(1)
}

func main() {
	LoadIgnore(".")
	c, err := client.NewClient()
	if err != nil {
		fatalError("Failed to create new API client", err)
	}
	localFiles, err := LocalAssetState(".")
	if err != nil {
		fatalError("Failed to determine local asset state", err)
	}
	remoteFiles, err := RemoteAssetState(c)
	if err != nil {
		fatalError("Failed to determine remote asset state", err)
	}
	diff := diffAssets(localFiles, remoteFiles)
	printDiff(diff)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	sync(c, diff)
}
