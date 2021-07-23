package client

// Sync performs all synchronization steps defined in an asset diff
// The callback function is called on completion of any actions
// If callback returns an error, the synchronization is halted
func (c *Client) Sync(diff []AssetDiff, callback func(Asset, error) error) {
	for _, i := range diff {
		var err error = nil
		switch i.Action {
		case Create, Update:
			err = callback(i.Asset, c.PutAsset(i.Asset.Path, i.Asset.Content))
		case Delete:
			err = callback(i.Asset, c.DeleteAsset(i.Asset.Path))
		}
		if err != nil {
			return
		}
	}
}
