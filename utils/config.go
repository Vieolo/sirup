package utils

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type RepoInConfig struct {
	Name     string `yaml:"name"`
	URL      string `yaml:"url"`
	RepoPath string `yaml:"path"`
}

type WorkspaceConfig struct {
	Name         string         `yaml:"name"`
	ProjectsPath string         `yaml:"projects_path"`
	Repos        []RepoInConfig `yaml:"repos"`
}

func ReadConfig() (WorkspaceConfig, error) {
	buf, err := os.ReadFile("sirup.workspace.yaml")
	if err != nil {
		return WorkspaceConfig{}, err
	}

	config := &WorkspaceConfig{
		ProjectsPath: "./projects",
	}
	err = yaml.Unmarshal(buf, config)
	if err != nil {
		return WorkspaceConfig{}, fmt.Errorf("in file %q: %w", "sirup.workspace.yaml", err)
	}

	return *config, err
}
