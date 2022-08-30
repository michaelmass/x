package marshal

import (
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/michaelmass/x/fs"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type dataEncoderFunc func(any) ([]byte, error)

var encodeDataTypeMap = map[string]dataEncoderFunc{
	"yaml": marshalYamlData,
	"yml":  marshalYamlData,
	"json": marshalJSONData,
}

// MarshalToFile encodes the data into a file
func MarshalToFile(data any, filename string) error {
	dir := filepath.Dir(filename)
	err := fs.MkdirAll(dir)

	if err != nil {
		return errors.Wrapf(err, "Error creating directories to file %s", filename)
	}

	content, err := Marshal(data, filepath.Ext(filename))

	if err != nil {
		return errors.Wrap(err, "Error encoding data")
	}

	err = fs.WriteFile(filename, content)

	if err != nil {
		return errors.Wrapf(err, "Error creating file %s", filename)
	}

	return nil
}

// Marshal encodes data into byte array
func Marshal(data any, datatype string) ([]byte, error) {
	dataMarshalFunc, ok := encodeDataTypeMap[strings.TrimPrefix(datatype, ".")]

	if !ok {
		return nil, errors.Errorf("Error encoding type %s isn't supported", datatype)
	}

	return dataMarshalFunc(data)
}

func marshalYamlData(data any) ([]byte, error) {
	return yaml.Marshal(data)
}

func marshalJSONData(data any) ([]byte, error) {
	return json.Marshal(data)
}
