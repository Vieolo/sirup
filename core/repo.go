package core

import (
	"slices"

	"github.com/vieolo/filange"
	"github.com/vieolo/termange"
)

type Repo struct {
	// Name of the repo
	Name string `yaml:"name"`
	// The URL of repo, including the https prefix
	URL string `yaml:"url"`
	// The relative path of the repo, in relation to the workspace root. e.g. frontend/my-project
	RepoPath string `yaml:"path"`
	// The type of the repo. e.g. `go`, `flutter`, etc.
	RepoType string `yaml:"type"`
	// Optional tags of the repo. Each repo can have many tags
	Tags []string `yaml:"tags"`
	// The absolute path of repo, calcuated while parsing the workspace yaml file
	AbsolutePath string
}

// Clones the repo from its git URL
func (r Repo) CloneFromGit() error {
	filange.CreateDirIfNotExists(r.RepoPath, 0777)
	_, _, err := termange.RunCommand(termange.CommandConfig{
		Command: "git",
		Args: []string{
			"clone",
			r.URL,
			r.RepoPath,
		},
	})
	return err
}

// Filters the repos by the given type
func (c WorkspaceConfig) FilterReposByType(repoType string) []Repo {
	filtered := []Repo{}

	for _, r := range c.Repos {
		if r.RepoType == repoType {
			filtered = append(filtered, r)
		}
	}

	return filtered
}

// Filters the repos by the given tag
func (c WorkspaceConfig) FilterReposByTag(tag string) []Repo {
	filtered := []Repo{}

	for _, r := range c.Repos {
		if slices.Contains(r.Tags, tag) {
			filtered = append(filtered, r)
		}
	}

	return filtered
}
