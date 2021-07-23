package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/moyen-blog/sync-dir/client"
)

// printDiff prints the staged actions required to synchronize local with remote files
func printDiff(diff []client.AssetDiff) {
	fmt.Printf("%d action(s) staged\n", len(diff))
	for _, i := range diff {
		switch i.Action {
		case client.Create:
			fmt.Printf("\033[32mCREATE\033[0m\t%s\n", i.Asset.Path)
		case client.Update:
			fmt.Printf("\033[33mUPDATE\033[0m\t%s\n", i.Asset.Path)
		case client.Delete:
			fmt.Printf("\033[31mDELETE\033[0m\t%s\n", i.Asset.Path)
		}
	}
}

// printProgress is a callback that we'll pass to client.Sync()
// It simply prints the results of an action but does not halt sync
func printProgress(asset client.Asset, err error) error {
	if err == nil {
		fmt.Printf("\033[32mSUCCESS\033[0m\t%s\n", asset.Path)
	}
	handleError("Failed to synchronize", err, false)
	return nil // Continue with sync regardless of errors
}

// askForConfirmation prompts the user to confirm a proposed action
func askForConfirmation(s string) bool {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s [y/n]: ", s)
		response, err := reader.ReadString('\n')
		handleError("Failed to read user input", err, true)
		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}

// handleError prints an error message of not nil
// If fatal is true, kills process with non-zero error code
func handleError(message string, err error, fatal bool) {
	if err != nil {
		fmt.Printf("\033[31mERROR\033[0m\t%s \033[2m%s\033[0m\n", message, err.Error())
		if fatal {
			os.Exit(1)
		}
	}
}
