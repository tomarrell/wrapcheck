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

// A file present in the directory named "analysis_skip" will cause the primary
// analysis tests to skip this directory due to needing explicit tests.
const skipfile = "analysistest_skip"

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

			dirPath, err := filepath.Abs(path.Join("./testdata", f.Name()))
			assert.NoError(t, err)

			// Check if the test is marked for skipping analysistest
			if _, err := os.Stat(path.Join(dirPath, skipfile)); err == nil {
				t.Logf("skipping test: %s\n", t.Name())
				t.Skip()
			}

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

func TestRegexpCompileFail(t *testing.T) {
	configFile, err := os.ReadFile("./testdata/config_ignoreSigRegexps_fail/.wrapcheck.yaml")
	assert.NoError(t, err)

	var config WrapcheckConfig
	assert.NoError(t, yaml.Unmarshal(configFile, &config))

	a := NewAnalyzer(config)

	results, err := a.Run(nil) // Doesn't matter what we pass
	assert.Nil(t, results)
	assert.Contains(t, err.Error(), "unable to compile regexp json\\.[a-zA-Z0-9_-")
}
