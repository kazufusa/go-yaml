package yaml

import (
	"errors"
	"testing"
)

// https://github.com/goccy/go-yaml/issues/273
func TestIssue273(t *testing.T) {
	source := "foo: |\n\ta b\n\tc\nbar: hello\n"
	var x struct {
		Foo string `yaml:"foo"`
		Bar string `yaml:"bar"`
	}
	actual := Unmarshal([]byte(source), &x)
	expect := errors.New("need new error")
	t.Fatalf("failed to test [%#v], actual=[%#v], expect=[%#v]", source, actual, expect)
}
