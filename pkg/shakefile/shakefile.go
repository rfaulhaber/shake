package shakefile

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type RunContext struct {
	Quiet bool
	Silent bool
}

type Shakefile struct {
	Targets map[string][]string
	Vars    map[string]string
	Default string
}

func DecodeFile(reader io.Reader) (Shakefile, error) {
	contents, err := ioutil.ReadAll(reader)

	if err != nil {
		return Shakefile{}, err
	}

	var sf Shakefile

	err = yaml.Unmarshal(contents, &sf)

	if err != nil {
		return Shakefile{}, err
	}

	return sf, nil
}

func (sf Shakefile) SetEnv() error {
	for key, value := range sf.Vars {
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	return nil
}

func (sf Shakefile) Run(target string, outWriter io.Writer, errorWriter io.Writer, context RunContext) error {
	targetCmds := sf.Targets[target]

	if targetCmds == nil {
		return errors.New(fmt.Sprintf(`no target called "%s"`, target))
	}

	steps := len(targetCmds)

	for cmdCount, commandLine := range targetCmds {
		command := strings.Split(commandLine, " ")
		args := command[1:]

		for index, arg := range args {
			if strings.HasPrefix(arg, "$") && sf.Vars[arg[1:]] != "" {
				args[index] = sf.Vars[arg[1:]]
			}
		}


		cmd := exec.Command(command[0], command[1:]...)
		// TODO: cmd.Env = os.Environ() + all sf.Vars fields

		if !context.Silent {
			cmd.Stdout = outWriter
			cmd.Stderr = errorWriter
		}

		if !context.Quiet && !context.Silent {
			fmt.Fprintf(outWriter, "%s %d/%d:  %s\n", target, cmdCount+1, steps, commandLine)
		}

		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}
