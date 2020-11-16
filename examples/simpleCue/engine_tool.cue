package example


import (
	"tool/exec"
)

command: engine: {

 task: run: exec.Run & {
  cmd: "cue eval -p example --out json --ignore --inject action=\(_action)"
 }
}
