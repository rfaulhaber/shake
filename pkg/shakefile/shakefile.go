package shakefile

import (
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

type Shakefile struct {
	Targets map[string][]string
	Vars    map[string]string
	Default string
	Requires []string
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

func (sf *Shakefile) expandVars() {
	// TODO implement
}


