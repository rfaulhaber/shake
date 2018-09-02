package shakefile

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"os"
	"testing"
)

func TestDecodeFile(t *testing.T) {
	expectedShakefile := makeTestFile()

	testBytes, err := yaml.Marshal(expectedShakefile)

	if err != nil {
		panic(err)
	}

	testBuf := bytes.NewReader(testBytes)

	shakefile, err := DecodeFile(testBuf)

	assert.NoError(t, err)

	assert.EqualValues(t, expectedShakefile, shakefile)
}

func TestShakefile_SetEnv(t *testing.T) {
	testShakefile := makeTestFile()

	testShakefile.SetEnv()

	for key, value := range testShakefile.Vars {
		assert.Equal(t, value, os.Getenv(key))
	}
}

func TestShakefile_Run(t *testing.T) {
	testCases := []struct {
		shakefile      Shakefile
		target         string
		expectedOutput string
		expectedErr    string
	}{
		{
			shakefile:      makeTestFile(),
			target:         "build",
			expectedOutput: "build",
			expectedErr:    "",
		},
		{
			shakefile:      makeTestFile(),
			target:         "release",
			expectedOutput: "release",
			expectedErr:    "",
		},
		{
			shakefile:      makeTestFile(),
			target:         "releaseWithEnv",
			expectedOutput: "build",
			expectedErr:    "",
		},
		{
			shakefile:      makeTestFile(),
			target:         "test",
			expectedOutput: "",
			expectedErr:    `no target called "test"`,
		},
		{
			shakefile:      makeTestFile(),
			target:         "targetWithoutEnv",
			expectedOutput: "$targetmsg",
			expectedErr:    "",
		},
	}

	for _, tc := range testCases {
		var outReader bytes.Buffer
		var errReader bytes.Buffer

		tc.shakefile.SetEnv()
		err := tc.shakefile.Run(tc.target, &outReader, &errReader)

		if err != nil {
			assert.Equal(t, err.Error(), tc.expectedErr)
		}

		assert.Equal(t, tc.expectedOutput, outReader.String())
	}
}

func makeTestFile() Shakefile {
	return Shakefile{
		Vars: map[string]string{
			"buildmsg":   "build",
			"releasemsg": "release",
		},
		Targets: map[string][]string{
			"build": {
				"echo -n build",
			},
			"release": {
				"echo -n release",
			},
			"releaseWithEnv": {
				"echo -n $buildmsg",
			},
			"targetWithoutEnv": {
				"echo -n $targetmsg",
			},
		},
	}
}
