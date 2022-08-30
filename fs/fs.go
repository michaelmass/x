package fs

import (
	"os"

	"github.com/pkg/errors"
)

const (
	dirModePerm  = 755
	fileModePerm = 644
)

func WriteFile(path string, content []byte) error {
	return os.WriteFile(path, content, fileModePerm)
}

func MkdirAll(path string) error {
	return os.MkdirAll(path, dirModePerm)
}

func Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func Copy(source string, destination string) error {
	content, err := os.ReadFile(source)

	if err != nil {
		return errors.Wrapf(err, "reading file %s", source)
	}

	err = os.WriteFile(destination, content, fileModePerm)

	if err != nil {
		return errors.Wrapf(err, "writing file %s", destination)
	}

	return nil
}
