package engine


import (
	"tool/exec"
)

command: engine: {

 task: run: exec.Run & {
  cmd: "cue eval --out json --ignore --inject action=\(_action)"
 }
}
