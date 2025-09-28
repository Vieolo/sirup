package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type RepoInConfig struct {
	Name     string   `yaml:"name"`
	URL      string   `yaml:"url"`
	RepoPath string   `yaml:"path"`
	RepoType string   `yaml:"type"`
	Tags     []string `yaml:"tags"`
}

type WorkspaceConfig struct {
	Name         string         `yaml:"name"`
	ProjectsPath string         `yaml:"projects_path,omitempty"`
	Repos        []RepoInConfig `yaml:"repos"`
}

func ReadConfig() (WorkspaceConfig, error) {
	buf, err := os.ReadFile("sirup.workspace.yaml")
	if err != nil {
		return WorkspaceConfig{}, err
	}

	config := &WorkspaceConfig{
		ProjectsPath: ".",
	}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		return WorkspaceConfig{}, fmt.Errorf("in file %q: %w", "sirup.workspace.yaml", err)
	}

	return *config, err
}

func WriteConfig(con WorkspaceConfig) error {
	if len(con.Repos) == 0 {
		con.Repos = append(con.Repos, RepoInConfig{
			Name:     "sample-repo",
			URL:      "https://samplegit.com/sample-repo",
			RepoPath: "allSamples/sample-repo",
			RepoType: "typescript",
			Tags:     []string{"frontend", "react", "vite"},
		})
	}

	marshaledBytes, marshalErr := yaml.Marshal(con)
	if marshalErr != nil {
		return marshalErr
	}

	return os.WriteFile("sirup.workspace.yaml", marshaledBytes, 0777)
}
