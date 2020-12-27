package main

import (
	"fmt"

	"github.com/moyen-blog/sync-dir/client"
)

func follow(progress chan error, done chan bool) {
	fmt.Printf("Executing %d action(s) ", cap(progress))
	for err := range progress {
		if err == nil {
			fmt.Print("\033[32m█\033[0m")
		} else {
			fmt.Print("\033[31m█\033[0m")
		}
	}
	fmt.Println(" Done!")
	done <- true
}

func sync(author string, token string, diff []AssetDiff) {
	progress, done := make(chan error, len(diff)), make(chan bool, 1)
	go follow(progress, done)
	for _, i := range diff {
		switch i.Action {
		case Create, Update:
			b, _ := i.Asset.Buffer() // Will be caught by client.PutAsset
			progress <- client.PutAsset(author, token, i.Asset.Path, b)
		case Delete:
			progress <- client.DeleteAsset(author, token, i.Asset.Path)
		}
	}
	close(progress)
	<-done
}
