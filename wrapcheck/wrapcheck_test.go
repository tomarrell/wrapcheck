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
