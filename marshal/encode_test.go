package marshal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func marshalTestingUtil(t *testing.T, value interface{}, expected string, datatype string) {
	content, err := Marshal(value, datatype)
	assert.Equal(t, expected, string(content))
	assert.Nil(t, err)
}

func TestMarshal(t *testing.T) {
	marshalTestingUtil(t, nilValue, nilValueExpectedJSON, "json")
	marshalTestingUtil(t, nilValue, nilValueExpectedYaml, "yaml")
	marshalTestingUtil(t, nilValue, nilValueExpectedYaml, "yml")
	marshalTestingUtil(t, nilValue, nilValueExpectedJSON, ".json")
	marshalTestingUtil(t, nilValue, nilValueExpectedYaml, ".yaml")
	marshalTestingUtil(t, nilValue, nilValueExpectedYaml, ".yml")

	marshalTestingUtil(t, intValue, intValueExpectedJSON, "json")
	marshalTestingUtil(t, intValue, intValueExpectedYaml, "yaml")
	marshalTestingUtil(t, intValue, intValueExpectedYaml, "yml")
	marshalTestingUtil(t, intValue, intValueExpectedJSON, ".json")
	marshalTestingUtil(t, intValue, intValueExpectedYaml, ".yaml")
	marshalTestingUtil(t, intValue, intValueExpectedYaml, ".yml")

	marshalTestingUtil(t, boolValue, boolValueExpectedJSON, "json")
	marshalTestingUtil(t, boolValue, boolValueExpectedYaml, "yaml")
	marshalTestingUtil(t, boolValue, boolValueExpectedYaml, "yml")
	marshalTestingUtil(t, boolValue, boolValueExpectedJSON, ".json")
	marshalTestingUtil(t, boolValue, boolValueExpectedYaml, ".yaml")
	marshalTestingUtil(t, boolValue, boolValueExpectedYaml, ".yml")

	marshalTestingUtil(t, mapString, mapStringExpectedJSON, "json")
	marshalTestingUtil(t, mapString, mapStringExpectedYaml, "yaml")
	marshalTestingUtil(t, mapString, mapStringExpectedYaml, "yml")
	marshalTestingUtil(t, mapString, mapStringExpectedJSON, ".json")
	marshalTestingUtil(t, mapString, mapStringExpectedYaml, ".yaml")
	marshalTestingUtil(t, mapString, mapStringExpectedYaml, ".yml")

	marshalTestingUtil(t, structValue, structValueExpectedJSON, "json")
	marshalTestingUtil(t, structValue, structValueExpectedYaml, "yaml")
	marshalTestingUtil(t, structValue, structValueExpectedYaml, "yml")
	marshalTestingUtil(t, structValue, structValueExpectedJSON, ".json")
	marshalTestingUtil(t, structValue, structValueExpectedYaml, ".yaml")
	marshalTestingUtil(t, structValue, structValueExpectedYaml, ".yml")
}
