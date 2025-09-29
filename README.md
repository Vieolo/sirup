# sirup
Multi-repo workspace manager, the closest thing to a monorepo without being a monorepo.

`sirup` can be used both as a CLI and a dependency for your Go application.

## Install as CLI
`sirup` is a standalone binary

```bash
go install github.com/vieolo/sirup@latest
```

## Use as dependency
Add `sirup` to your project

```bash
go get github.com/vieolo/sirup
```

You can import the functionalities from `sirup/core`

```go
import (
  sirup "github.com/vieolo/sirup/core"
)

func main() {
  config, configErr := sirup.ReadWorkspaceConfig()
}
```

## Workspace
A sirup workspace is a folder that contains the `sirup.workspace.yaml` config file. It is recommended to create a standalone repo to hold a workspace.

Each workspace can have multiple independent repos, listed in the `sirup.workspace.yaml` which will be cloned in the given path in the workspace.

Here is the structure of the `sirup.workspace.yaml` file. I suggest avoid using whitespaces in the names.

```yaml
name: "workspace_name" # Name of the workspace -- Required
projects_path: "." # Where to clone the projects -- Optional. defaults: "."

repos: # The array of repos. Should have at least one repo -- Required
    - name: "repo_1" # Name of the repo. can be different than the actual repo name -- Required
      url: "https://github.com/to/your/repo_1" # URL of the repo -- Required
      path: "utils/repo_1" # Where to clone the repo inside the workspace -- Required
      type: "frontend" # The type of the repo -- Optional
      tags:
        - typescript
        - vite
        - react
    - name: "repo_2"
      url: "https://github.com/to/your/repo_2"
      path: "cicd/repo_2"
      type: "go"
```

## Commands

- `init` -> Initiates a new workspace. Will prompt you for the necessary fields
- `fetch` -> Clones the repos listed in workspace yaml
- `list` -> Lists the repos listed in workspace yaml
