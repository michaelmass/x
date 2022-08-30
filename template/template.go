package template

import (
	"io"
	templ "text/template"

	"github.com/Masterminds/sprig"
)

var funcMap = sprig.TxtFuncMap()

func Render(content []byte, data any, wr io.Writer) error {
	t, err := templ.New("default").Funcs(funcMap).Parse(string(content))

	if err != nil {
		return err
	}

	t.Option("missingkey=error")

	err = t.Execute(wr, data)

	if err != nil {
		return err
	}

	return nil
}

func RenderFile(reader io.Reader, data any, wr io.Writer) error {
	content, err := io.ReadAll(reader)

	if err != nil {
		return err
	}

	return Render(content, data, wr)
}
