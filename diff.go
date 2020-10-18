package main

import "fmt"

// Action defines the Markdown file action enum underlying type
type Action uint

// Markdown file action enum
const (
	Create Action = iota
	Update
	Delete
)

// MarkdownDiff defines an action that should be taken on a Markdown file
type MarkdownDiff struct {
	File   MarkdownFile
	Action Action
}

func sliceContains(a MarkdownFile, b []MarkdownFile) (bool, *MarkdownFile) {
	for _, i := range b {
		if a.Path == i.Path {
			return true, &i
		}
	}
	return false, nil
}

func diffMarkdownFiles(localFiles []MarkdownFile, remoteFiles []MarkdownFile) (diff []MarkdownDiff) {
	for _, local := range localFiles { // Find ∈a & ∉b and ∈a & ∈b
		if found, remoteFile := sliceContains(local, remoteFiles); found {
			if local.Hash != remoteFile.Hash {
				fmt.Println("Update", local.Path)
				diff = append(diff, MarkdownDiff{
					File:   local,
					Action: Update,
				})
			}
		} else {
			fmt.Println("Create", local.Path)
			diff = append(diff, MarkdownDiff{
				File:   local,
				Action: Create,
			})
		}
	}
	for _, remote := range remoteFiles { // Find ∈b & ∉a
		if found, _ := sliceContains(remote, localFiles); !found {
			fmt.Println("Delete", remote.Path)
			diff = append(diff, MarkdownDiff{
				File:   remote,
				Action: Delete,
			})
		}
	}
	return
}
