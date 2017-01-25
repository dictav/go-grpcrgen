package grpcrgen

import (
	"bytes"
	"io/ioutil"
	"testing"
	"text/template"
)

func TestHelperTemplate(t *testing.T) {
	st := struct {
		PackageName string
	}{"test"}

	b := make([]byte, 1024)
	w := bytes.NewBuffer(b)

	buf, err := ioutil.ReadFile("template/helper.tmpl")
	if err != nil {
		t.Error(err)
	}

	tmpl := template.New("t")
	tmpl.Parse(string(buf))
	if err := tmpl.Execute(w, st); err != nil {
		t.Error(err)
	}
}

func TestRouterTemplate(t *testing.T) {
	ret := struct {
		PackageName string
		API         []api
	}{
		"test_router",
		[]api{{"service", "import_path", []string{"funcA", "funcB"}}},
	}

	b := make([]byte, 1024)
	w := bytes.NewBuffer(b)

	buf, err := ioutil.ReadFile("template/router.tmpl")
	if err != nil {
		t.Error(err)
	}

	tmpl := template.New("t")
	tmpl.Parse(string(buf))
	if err := tmpl.Execute(w, ret); err != nil {
		t.Error(err)
	}
}
