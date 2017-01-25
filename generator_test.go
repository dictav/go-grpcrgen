package grpcrgen

import (
	"bytes"
	"regexp"
	"strings"
	"testing"
)

func TestGenerateHelper(t *testing.T) {
	buf := make([]byte, 0, 2048)
	w := bytes.NewBuffer(buf)

	err := generateHelper(w, "testpackage")
	if err != nil {
		t.Error(err)
	}

	out := string(w.Bytes())
	lines := strings.Split(out, "\n")

	want := "package testpackage"
	got := lines[0]
	if want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}
}

func TestGenerateRouter(t *testing.T) {
	buf := make([]byte, 0, 4096)
	w := bytes.NewBuffer(buf)

	err := generateRouter(w, "test_data", "testpackage")
	if err != nil {
		t.Error(err)
	}

	out := string(w.Bytes())
	lines := strings.Split(out, "\n")

	want := "package testpackage"
	got := lines[0]
	if want != got {
		t.Errorf("want '%s', got '%s'", want, got)
	}

	test1Router := regexp.MustCompile(`func NewTest1ServiceRouter`)
	test2Router := regexp.MustCompile(`func NewTest2ServiceRouter\(`)
	testNRouter := regexp.MustCompile(`func NewTestNestServiceRouter\(`)

	if ret := test1Router.FindAll((w.Bytes()), 10); len(ret) != 1 {
		t.Errorf("want 1, got %d", len(ret))
	}

	if ret := test2Router.FindAll(w.Bytes(), 10); len(ret) != 1 {
		t.Errorf("want 1, got %d", len(ret))
	}

	if ret := testNRouter.FindAll(w.Bytes(), 10); len(ret) != 1 {
		t.Errorf("want 1, got %d", len(ret))
	}
}

func TestExtractAPI(t *testing.T) {
	apis, err := extractAPI("test_data")
	if err != nil {
		t.Error(err)
	}

	if len(apis) != 3 {
		t.Errorf("want 3, got %d", len(apis))
	}
}
