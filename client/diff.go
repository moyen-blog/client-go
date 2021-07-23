package client

// Action defines the file action enum underlying type
type Action int

// File action enum
const (
	Create Action = iota
	Update
	Delete
)

// AssetDiff defines an action that should be taken on a file
type AssetDiff struct {
	Asset  Asset
	Action Action
}

func sliceContains(a Asset, b []Asset) (bool, *Asset) {
	for _, i := range b {
		if a.Path == i.Path {
			return true, &i
		}
	}
	return false, nil
}

// DiffAssets compares local and remote assets and returns a diff
func (c *Client) DiffAssets(localFiles []Asset, remoteFiles []Asset) (diff []AssetDiff) {
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
