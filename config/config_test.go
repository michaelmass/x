package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ConfigTest struct {
	Name string `yaml:"name"`
}

func NewConfigTest() *ConfigTest {
	return &ConfigTest{
		Name: "toolname",
	}
}

func TestConfig(t *testing.T) {
	configClient := New("toolname", NewConfigTest)
	home, _ := os.UserHomeDir()

	dir, err := configClient.Dir()
	assert.NoError(t, err)
	assert.Equal(t, filepath.Join(home, "toolname"), dir)

	path, err := configClient.Path()
	assert.NoError(t, err)
	assert.Equal(t, filepath.Join(home, "toolname/config.yml"), path)
}
