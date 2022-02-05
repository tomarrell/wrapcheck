package wrapcheck

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"
	"gopkg.in/yaml.v3"
)

func TestAnalyzer(t *testing.T) {
	// Load the dirs under ./testdata
	p, err := filepath.Abs("./testdata")
	assert.NoError(t, err)

	files, err := ioutil.ReadDir(p)
	assert.NoError(t, err)

	for _, f := range files {
		t.Run(f.Name(), func(t *testing.T) {
			if !f.IsDir() {
				t.Fatalf("cannot run on non-directory: %s", f.Name())
			}

			if f.Name() == "config_ignoreSigRegexps_fail" {
				t.Skipf("skipping %s, as it expect to fail in this test", f.Name())
			}

			dirPath, err := filepath.Abs(path.Join("./testdata", f.Name()))
			assert.NoError(t, err)

			configPath := path.Join(dirPath, ".wrapcheck.yaml")
			if _, err := os.Stat(configPath); os.IsNotExist(err) {
				// There is no config
				analysistest.Run(t, dirPath, NewAnalyzer(NewDefaultConfig()))
			} else if err == nil {
				// A config file exists, use it
				configFile, err := os.ReadFile(configPath)
				assert.NoError(t, err)

				var config WrapcheckConfig
				assert.NoError(t, yaml.Unmarshal(configFile, &config))
				analysistest.Run(t, dirPath, NewAnalyzer(config))

			} else {
				assert.FailNow(t, err.Error())
			}
		})
	}
}

func TestFail(t *testing.T) {
	// A config file exists, use it
	configFile, err := os.ReadFile("./testdata/config_ignoreSigRegexps_fail/.wrapcheck.yaml")
	assert.NoError(t, err)

	var config WrapcheckConfig
	assert.NoError(t, yaml.Unmarshal(configFile, &config))
	a := NewAnalyzer(config)
	results, err := a.Run(nil) // doesn't matter what we passing ...

	assert.Nil(t, results)
	assert.EqualError(t, err,
		"unable to parse regexp: error parsing regexp: missing closing ]: `[a-zA-Z0-9_-` at json\\.[a-zA-Z0-9_-\n")
}
