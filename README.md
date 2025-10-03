# sirup
Multi-repo workspace manager, the closest thing to a monorepo without being a monorepo.

`sirup` can be used both as a CLI and a dependency for your Go application.


## Features

- Create a workspace with a yaml config file
- Cloning and listing all repos of the workspace with single command


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

## Topics

- [Workspaces and their config file](./doc/workspace.md)
- [CLI commands](./doc/cli.md)
