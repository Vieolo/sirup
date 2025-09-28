package utils

import (
	"fmt"
	"os"
	"path/filepath"

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
	configPath, pathErr := findWorkspaceFile()
	if pathErr != nil {
		return WorkspaceConfig{}, pathErr
	}
	buf, err := os.ReadFile(configPath)
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

// it searches for the workspace file
// starting in the current directory and continuing up through parent directories
// until it reaches the filesystem root.
func findWorkspaceFile() (string, error) {
	// 1. Get the current working directory to start the search.
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current working directory: %w", err)
	}

	filename := "sirup.workspace.yaml"

	// Loop indefinitely, breaking only when the file is found or the root is reached.
	for {
		// 2. Construct the full path to the potential file.
		filePath := filepath.Join(currentDir, filename)

		// 3. Check if the file exists at the current path.
		_, err := os.Stat(filePath)

		if err == nil {
			// Success: File found. Return the full path.
			return filePath, nil
		}

		// 4. If the error is not "file does not exist", something else is wrong.
		// We stop here if it's a permission error or similar.
		if !os.IsNotExist(err) {
			return "", fmt.Errorf("error checking file existence at %s: %w", filePath, err)
		}

		// 5. Determine the parent directory.
		parentDir := filepath.Dir(currentDir)

		// 6. Check if we've reached the filesystem root.
		// On most OSs, `filepath.Dir("/")` returns "/", and `filepath.Dir("C:\\")` returns "C:\\".
		if parentDir == currentDir {
			// We are at the root and haven't found the file.
			return "", fmt.Errorf("file '%s' not found in current directory or any parent directories", filename)
		}

		// 7. Move up to the parent directory for the next iteration.
		currentDir = parentDir
	}
}
