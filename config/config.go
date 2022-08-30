package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/michaelmass/x/fs"
	"github.com/pkg/errors"
	"github.com/skratchdot/open-golang/open"
	"gopkg.in/yaml.v2"
)

const (
	configFilename = "config.yml"
)

type NewFunc[C any] func() *C

type Client[C any] struct {
	path string
	new  NewFunc[C]
}

func New[C any](path string, newFunc NewFunc[C]) *Client[C] {
	return &Client[C]{
		path: path,
		new:  newFunc,
	}
}

func (client *Client[C]) Dir() (string, error) {
	home, err := os.UserHomeDir()

	if err != nil {
		return "", errors.Wrap(err, "getting user home dir")
	}

	return filepath.Join(home, client.path), nil
}

func (client *Client[C]) Path() (string, error) {
	dir, err := client.Dir()

	if err != nil {
		return "", errors.Wrap(err, "getting config dir")
	}

	return filepath.Join(dir, configFilename), nil
}

func (client *Client[C]) Reset() error {
	path, err := client.Path()

	if err != nil {
		return errors.Wrap(err, "getting config path")
	}

	err = fs.MkdirAll(filepath.Dir(path))

	if err != nil {
		return errors.Wrap(err, "creating config folder")
	}

	config := client.new()
	content, err := yaml.Marshal(config)

	if err != nil {
		return errors.Wrap(err, "encoding yaml config file")
	}

	err = fs.WriteFile(path, content)

	if err != nil {
		return errors.Wrap(err, "writing config file")
	}

	return nil
}

func (client *Client[C]) Init() error {
	path, err := client.Path()

	if err != nil {
		return errors.Wrap(err, "getting config path")
	}

	if fs.Exist(path) {
		return nil
	}

	err = client.Reset()

	if err != nil {
		return errors.Wrap(err, "initializing config file")
	}

	return nil
}

func (client *Client[C]) FromFile(path string) (*C, error) {
	config := client.new()
	content, err := os.ReadFile(path)

	if err != nil {
		return nil, errors.Wrap(err, "reading config file")
	}

	err = yaml.Unmarshal(content, config)

	if err != nil {
		return nil, errors.Wrap(err, "decoding config file content")
	}

	return config, nil
}

func (client *Client[C]) Open() error {
	path, err := client.Path()

	if err != nil {
		return errors.Wrap(err, "getting config path")
	}

	return open.Run(path)
}

func (client *Client[C]) Show() error {
	path, err := client.Path()

	if err != nil {
		return errors.Wrap(err, "getting config path")
	}

	config, err := client.FromFile(path)

	if err != nil {
		return errors.Wrap(err, "reading config file")
	}

	content, err := yaml.Marshal(config)

	if err != nil {
		return errors.Wrap(err, "decoding content of config file")
	}

	fmt.Println(string(content))

	return nil
}

func (client *Client[C]) Get() (*C, error) {
	path, err := client.Path()

	if err != nil {
		return nil, errors.Wrap(err, "getting config path")
	}

	config, err := client.FromFile(path)

	if err != nil {
		return nil, errors.Wrap(err, "reading config file")
	}

	return config, nil
}
