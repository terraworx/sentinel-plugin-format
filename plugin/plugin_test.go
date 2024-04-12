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
			"Basic usage",
			`main = subject.string("Hello, %s!", ["world"]) == "Hello, world!"`,
		},
		{
			"Format specifiers",
			`main = subject.string("The value of pi is %.2f", [3.14159]) == "The value of pi is 3.14"`,
		},
		{
			"Multiple arguments",
			`main = subject.string("Values: %d, %s, %.2f", [42, "foo", 3.14]) == "Values: 42, foo, 3.14"`,
		},
		{
			"Named format verbs",
			`main = subject.string("The temperature is %d°F", [32]) == "The temperature is 32°F"`,
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
