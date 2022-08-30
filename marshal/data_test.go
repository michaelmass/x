package marshal

const (
	nilValueExpectedJSON = "null"
	nilValueExpectedYaml = "null\n"

	intValueExpectedJSON = "1"
	intValueExpectedYaml = "1\n"

	boolValueExpectedJSON = "false"
	boolValueExpectedYaml = "false\n"

	mapStringExpectedJSON = "{\"\":\"\",\"empty\":\"\",\"key\":\"value\",\"test\":\"test\"}"
	mapStringExpectedYaml = "\"\": \"\"\nempty: \"\"\nkey: value\ntest: test\n"

	structValueExpectedJSON = "{\"Test\":\"testing\",\"Value\":123}"
	structValueExpectedYaml = "test: testing\nvalue: 123\n"
)

var (
	nilValue interface{} = nil

	intValue = 1

	boolValue = false

	mapString = map[string]string{
		"test":  "test",
		"key":   "value",
		"empty": "",
		"":      "",
	}

	structValue = struct {
		Test       string
		Value      int
		unexported bool
	}{
		Test:       "testing",
		Value:      123,
		unexported: false,
	}

	mapStringExpectedJSONDecode = map[string]any(map[string]any{"": "", "empty": "", "key": "value", "test": "test"})
	mapStringExpectedYamlDecode = map[any]any(map[any]any{"": "", "empty": "", "key": "value", "test": "test"})
)
