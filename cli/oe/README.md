# OpenEgnine CLI
## Getting started
### Installation
OpenEngine CLI can be downloaded, installed or build:
  - Download the latest CLI binary release from GitHub at https://github.com/abraverm/openengine/releases
  - Install using Go `go get github.com/abraverm/openengine/cli/oe`
    

  - Build the CLI:
    ```bash
    git clone https://github.com/abraverm/openengine
    cd openengine
    go build ./cli/oe
    ```
    

    Note: Go version 1.16 is required to install or build OpenEngine

## Combination File

OpenEngine CLI creates a new instance of an engine and loads the different providers, provisioners, system and resources
specified in a Yaml file, a "combination file":

```yaml
---
providers: []
provisioners: []
systems: []
resources: []
```

All parameters expect a **list** of files locations. Location of files could be local (relative to dsl file or absolute)
or remote Http URL, but must be readable. All files must be Cue files. The Cue definition in the file must match the
parameter, i.e file with resource definition would be listed in `resources` parameter. File can contain only one
definition. The definition must have a valid Cue syntax (version 0.3.0-beta.5) and
match [Open Engine specifications](../../engine/spec.cue).