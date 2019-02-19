package depgen

import (
	"os/exec"

	"github.com/gobuffalo/genny"
	"github.com/pkg/errors"
)

func Ensure(verbose bool) (*genny.Generator, error) {
	g := genny.New()

	var args []string
	if verbose {
		args = append(args, "-v")
	}

	id, err := InstallDep(args...)
	if err != nil {
		return g, errors.WithStack(err)
	}
	g.Merge(id)

	cmd := exec.Command("dep", "ensure")
	if verbose {
		cmd.Args = append(cmd.Args, args...)
	}
	g.Command(cmd)
	return g, nil
}
