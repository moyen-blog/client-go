package main

import (
	"github.com/moyen-blog/sync-dir/asset"
)

// Action defines the file action enum underlying type
type Action uint

// File action enum
const (
	Create Action = iota
	Update
	Delete
)

// AssetDiff defines an action that should be taken on a file
type AssetDiff struct {
	Asset  asset.Asset
	Action Action
}

func sliceContains(a asset.Asset, b []asset.Asset) (bool, *asset.Asset) {
	for _, i := range b {
		if a.Path == i.Path {
			return true, &i
		}
	}
	return false, nil
}

func diffAssets(localFiles []asset.Asset, remoteFiles []asset.Asset) (diff []AssetDiff) {
	for _, local := range localFiles { // Find ∈a & ∉b and ∈a & ∈b
		if found, remoteFile := sliceContains(local, remoteFiles); found {
			if local.Hash != remoteFile.Hash {
				diff = append(diff, AssetDiff{
					Asset:  local,
					Action: Update,
				})
			}
		} else {
			diff = append(diff, AssetDiff{
				Asset:  local,
				Action: Create,
			})
		}
	}
	for _, remote := range remoteFiles { // Find ∈b & ∉a
		if found, _ := sliceContains(remote, localFiles); !found {
			diff = append(diff, AssetDiff{
				Asset:  remote,
				Action: Delete,
			})
		}
	}
	return
}
