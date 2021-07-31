package main

import (
	"fmt"
	"testing"
	"testing/fstest"
)

func TestConfigEmptyYAML(t *testing.T) {
	path := "test.yaml"
	dest := struct{}{}
	fsys := make(fstest.MapFS)
	fsys[path] = &fstest.MapFile{
		Data: []byte(""),
	}
	err := parseYaml(fsys, path, true, &dest)
	if err != nil {
		t.Error("Should successfully parse config YAML")
	}
}

func TestConfigMixedYAML(t *testing.T) {
	path := "test.yaml"
	yaml := `int: 1
string: string
array:
  - 1`
	dest := struct {
		Int    int
		String string
		Array  []int
	}{}
	fsys := make(fstest.MapFS)
	fsys[path] = &fstest.MapFile{
		Data: []byte(yaml),
	}
	err := parseYaml(fsys, path, true, &dest)
	if err != nil {
		t.Error("Should successfully parse config YAML")
	}
	if dest.Int != 1 {
		t.Errorf("Should parse int value of %d but got %d", 1, dest.Int)
	}
	if dest.String != "string" {
		t.Errorf("Should parse string value of %s but got %s", "string", dest.String)
	}
	if len(dest.Array) != 1 {
		t.Errorf("Should parse int array of length %d but got %d", 1, len(dest.Array))
	}
	if dest.Array[0] != 1 {
		t.Errorf("Should parse int array with member %d but got %d", 1, dest.Array[0])
	}
}

func TestConfigMissingRequiredYAML(t *testing.T) {
	path := "test.yaml"
	dest := struct{}{}
	fsys := make(fstest.MapFS)
	err := parseYaml(fsys, path, true, &dest)
	if err == nil {
		t.Error("Should fail if missing required YAML")
	}
}

func TestConfigMissingOptionalYAML(t *testing.T) {
	path := "test.yaml"
	fsys := make(fstest.MapFS)
	err := parseYaml(fsys, path, false, nil)
	if err != nil {
		t.Error("Should succeed if missing optional YAML")
	}
}

func TestConfigErrorBadYAML(t *testing.T) {
	path := "test.yaml"
	dest := struct{}{}
	fsys := make(fstest.MapFS)
	fsys[path] = &fstest.MapFile{
		Data: []byte("(*&^"),
	}
	err := parseYaml(fsys, path, true, &dest)
	if err == nil {
		t.Error("Should fail parsing invalid YAML")
	}
}

func TestConfigMoyenYAML(t *testing.T) {
	endpoint := "https://testendpoint.com"
	ignore := "test_*.md"
	username := "testusername"
	token := "testtoken"
	configYaml := fmt.Sprintf(`endpoint: %s
ignore:
  - %s`, endpoint, ignore)
	credentialsYaml := fmt.Sprintf(`username: %s
token: %s`, username, token)
	fsys := make(fstest.MapFS)
	fsys[".moyenconfig"] = &fstest.MapFile{
		Data: []byte(configYaml),
	}
	fsys[".moyencredentials"] = &fstest.MapFile{
		Data: []byte(credentialsYaml),
	}
	c, err := parseConfig(fsys)
	if err != nil {
		t.Error("Should successfully parse config and credentials YAML")
	}
	if c.credentials.Username != username {
		t.Errorf("Should have username %s but got %s", username, c.credentials.Username)
	}
	if c.credentials.Token != token {
		t.Errorf("Should have token %s but got %s", token, c.credentials.Token)
	}
	if c.Endpoint != endpoint {
		t.Errorf("Should have endpoint %s but got %s", endpoint, c.Endpoint)
	}
	if len(c.Ignore) != 2 {
		t.Errorf("Should ignore %d patterns but got %d", 2, len(c.Ignore))
	}
}
