package engine


import (
	"tool/exec"
    "tool/os"
)

command: engine: {
 task: run: exec.Run & {
  cmd: "cue eval --out \(env.format) --ignore --inject action=\(env.action)"
  env: os.Getenv & {
    format: *"yaml" | string
    action: *"create" | string
  }
 }
}
