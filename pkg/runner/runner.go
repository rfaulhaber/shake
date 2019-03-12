package runner

import (
	"io"
	"shake/pkg/shakefile"
)

type context struct {

}

type Runner struct {
	Stdout io.Writer
	Stdin io.Reader
	Stderr io.Writer

	sf shakefile.Shakefile
}

func (r *Runner) RunTarget(target string) {
}
