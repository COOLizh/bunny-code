package executor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gitlab.com/greenteam1/executioner/pkg/models"
)

// BuildConfig struct contains build configs
type BuildConfig struct {
	Config map[string]*models.BuildConfig
}

// NewBuildConfig creates new instance of BuildConfig
func NewBuildConfig() *BuildConfig {
	return &BuildConfig{}
}

// Read reads build configs and writes it to BuildConfig
func (conf *BuildConfig) Read(path string) error {
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	defer func() {
		closeErr := f.Close()
		if closeErr != nil {
			log.Println(err)
		}
	}()
	val, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	err = json.Unmarshal(val, &conf.Config)
	return err
}
