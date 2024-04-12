package format

import (
	"os"
	"testing"

	sdk "github.com/hashicorp/sentinel-sdk"
	framework "github.com/hashicorp/sentinel-sdk/framework"
	plugintesting "github.com/hashicorp/sentinel-sdk/testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	plugintesting.Clean()
	os.Exit(exitCode)
}

func TestImport_impl(t *testing.T) {
	var _ sdk.Plugin = New()
}

func TestRoot_impl(t *testing.T) {
	var _ framework.Root = new(root)
}

func TestImport(t *testing.T) {

	cases := []struct {
		Name   string
		Source string
	}{
		{
			"string",
			`main = subject.string("This is a formatted string with the value %s", "foo") is "This is a formatted string with the value foo"`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			plugintesting.TestPlugin(t, plugintesting.TestPluginCase{
				Source: tc.Source,
			})
		})
	}
}
