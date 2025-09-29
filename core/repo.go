package core

import (
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
