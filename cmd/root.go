package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rfaulhaber/shake/pkg/shakefile"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	stderr = log.New(os.Stderr, "shake: ", 0)
)

var rootCmd = &cobra.Command{
	Use:   "shake [OPTIONS] [TARGETS]",
	Short: "Automated target-based building and deployment system",
	Long:  `Automated target-based building and deployment system`,
	Run: runShake,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runShake(cmd *cobra.Command, args []string) {
	file, err := readShakefile()

	if err != nil {
		stderr.Fatalln(err)
	}

	if len(args) == 0 {
		if file.Default == "" && len(file.Targets) == 1 {
			if len(file.Targets) == 1 {
				for target := range file.Targets {
					if err = runTarget(file, target); err != nil {
						stderr.Fatalln(err)
					}
				}
			} else {
				stderr.Fatalln(`no "default" found, more than one target specified. exiting`)
			}
		} else {
			if err = runTarget(file, file.Default); err != nil {
				stderr.Fatalln(err)
			}
		}
	} else {
		for _, target := range args {
			if err = runTarget(file, target); err != nil {
				stderr.Fatalln(err)
			}
		}
	}
}

func readShakefile() (shakefile.Shakefile, error) {
	infos, err := ioutil.ReadDir(".")

	var shakefileName string

	for _, info := range infos {
		filename := info.Name()
		ext := filepath.Ext(filename)
		name := filename[0:len(filename)-len(ext)]

		if name == "Shakefile" {
			shakefileName = filename
			break
		}
	}

	if shakefileName == "" {
		return shakefile.Shakefile{}, errors.New("no shakefile found")
	}

	file, err := os.Open(shakefileName)

	if err != nil && os.IsNotExist(err) {
		return shakefile.Shakefile{}, errors.New("no shakefile found")
	}

	return shakefile.DecodeFile(file)
}

func runTarget(file shakefile.Shakefile, target string) error {
	if err := file.Run(target, os.Stdout, os.Stderr); err != nil {
		return errors.Wrap(err, fmt.Sprintf("running target %s", target))
	}

	return nil
}
