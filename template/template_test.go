package template

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	templateEmpty   = ""
	template        = "this is a test"
	templateData    = "{{ .value }}"
	templateTxtFunc = "{{ \"hello!\" | upper | repeat 5 }}"

	templateTxtFuncExpected = "HELLO!HELLO!HELLO!HELLO!HELLO!"
)

var data = map[string]string{"value": "test"}

func TestRender(t *testing.T) {
	var b bytes.Buffer
	err := Render([]byte(templateEmpty), nil, &b)
	assert.Equal(t, templateEmpty, b.String())
	assert.Nil(t, err)
	b.Reset()

	err = Render([]byte(template), nil, &b)
	assert.Equal(t, template, b.String())
	assert.Nil(t, err)
	b.Reset()

	err = Render([]byte(templateData), data, &b)
	assert.Equal(t, "test", b.String())
	assert.Nil(t, err)
	b.Reset()

	err = Render([]byte(templateTxtFunc), nil, &b)
	assert.Equal(t, templateTxtFuncExpected, b.String())
	assert.Nil(t, err)
	b.Reset()

	err = Render([]byte(templateData), nil, &b)
	assert.Equal(t, "", b.String())
	assert.NotNil(t, err)
	b.Reset()
}

func TestRenderFile(t *testing.T) {
	var b bytes.Buffer
	err := RenderFile(strings.NewReader(templateEmpty), nil, &b)
	assert.Equal(t, templateEmpty, b.String())
	assert.Nil(t, err)
	b.Reset()

	err = RenderFile(strings.NewReader(template), nil, &b)
	assert.Equal(t, template, b.String())
	assert.Nil(t, err)
	b.Reset()

	err = RenderFile(strings.NewReader(templateData), data, &b)
	assert.Equal(t, "test", b.String())
	assert.Nil(t, err)
	b.Reset()

	err = RenderFile(strings.NewReader(templateTxtFunc), nil, &b)
	assert.Equal(t, templateTxtFuncExpected, b.String())
	assert.Nil(t, err)
	b.Reset()

	err = RenderFile(strings.NewReader(templateData), nil, &b)
	assert.Equal(t, "", b.String())
	assert.NotNil(t, err)
	b.Reset()
}
