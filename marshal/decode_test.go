package marshal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func unmarshalTestingUtil(t *testing.T, value string, expected any, datatype string) {
	var data any
	err := Unmarshal([]byte(value), datatype, &data)
	assert.Equal(t, expected, data)
	assert.Nil(t, err)
}

func TestDecode(t *testing.T) {
	unmarshalTestingUtil(t, mapStringExpectedJSON, mapStringExpectedJSONDecode, "json")
	unmarshalTestingUtil(t, mapStringExpectedYaml, mapStringExpectedYamlDecode, "yaml")
	unmarshalTestingUtil(t, mapStringExpectedJSON, mapStringExpectedJSONDecode, ".json")
	unmarshalTestingUtil(t, mapStringExpectedYaml, mapStringExpectedYamlDecode, ".yaml")

	unmarshalTestingUtil(t, boolValueExpectedJSON, boolValue, "json")
	unmarshalTestingUtil(t, boolValueExpectedYaml, boolValue, "yaml")
	unmarshalTestingUtil(t, boolValueExpectedJSON, boolValue, ".json")
	unmarshalTestingUtil(t, boolValueExpectedYaml, boolValue, ".yaml")
}
