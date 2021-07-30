package client

import "testing"

func TestDiffIdentical(t *testing.T) {
	username := "testuser"
	token := "testtoken"
	c, err := NewClient(username, token, "", []string{})
	if err != nil {
		t.Error("Should create new client with valid credentials")
	}
	assets := make([]Asset, 0)
	diff := c.DiffAssets(assets, assets)
	if len(diff) != 0 {
		t.Errorf("Should produce a diff of %d assets but found %d", 0, len(diff))
	}
}

func TestDiffMixed(t *testing.T) {
	username := "testuser"
	token := "testtoken"
	c, err := NewClient(username, token, "", []string{})
	if err != nil {
		t.Error("Should create new client with valid credentials")
	}
	assetsLocal := []Asset{
		{Path: "testLocal.md"},
		{Path: "testIntersection.md", Hash: "A"},
	}
	assetsRemote := []Asset{
		{Path: "testRemote.md"},
		{Path: "testIntersection.md", Hash: "B"},
	}
	diff := c.DiffAssets(assetsLocal, assetsRemote)
	if len(diff) != 3 {
		t.Errorf("Should produce a diff of %d assets but found %d", 3, len(diff))
	}
	counts := make(map[Action]int)
	for _, d := range diff {
		counts[d.Action] += 1
	}
	if counts[Create] != 1 {
		t.Errorf("Should require %d creates but found %d", 1, counts[Create])
	}
	if counts[Update] != 1 {
		t.Errorf("Should require %d updates but found %d", 1, counts[Update])
	}
	if counts[Delete] != 1 {
		t.Errorf("Should require %d deletes but found %d", 1, counts[Delete])
	}
}
