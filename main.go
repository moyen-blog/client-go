package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/moyen-blog/sync-dir/client"
)

func main() {
	yes := flag.Bool("y", false, "skips all interactive prompts")
	flag.Parse()

	fsys := os.DirFS(".") // CWD

	config, err := parseConfig(fsys)
	handleError("Failed to load configuration JSON", err, true)

	c, err := client.NewClient(config.credentials.Username, config.credentials.Token, config.Endpoint, config.Ignore)
	handleError("Failed to create new API client", err, true)

	localFiles, err := c.AssetStateLocal(fsys)
	handleError("Failed to determine local asset state", err, true)

	remoteFiles, err := c.AssetStateRemote()
	handleError("Failed to determine remote asset state", err, true)

	diff := c.DiffAssets(localFiles, remoteFiles)
	if len(diff) == 0 {
		fmt.Println("No changes to sync")
		return
	}
	printDiff(diff)

	if !*yes { // Yes flag bypasses prompt
		if c := askForConfirmation("Proceed with changes?"); !c {
			return
		}
	}
	c.Sync(diff, printProgress)
}
