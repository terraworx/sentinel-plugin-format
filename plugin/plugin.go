package format

import (
	"fmt"

	sdk "github.com/hashicorp/sentinel-sdk"
	framework "github.com/hashicorp/sentinel-sdk/framework"
)

func New() sdk.Plugin {
	return &framework.Plugin{
		Root: &root{},
	}
}

type root struct {
}

func (m *root) Configure(raw map[string]interface{}) error {
	return nil
}

func (m *root) Get(key string) (interface{}, error) {
	return nil, nil
}

func (m *root) Func(key string) interface{} {
	switch key {
	case "string":
		return m.fmtString
	}
	return nil
}

func (m *root) fmtString(input string, args ...any) string {
	return fmt.Sprintf(input, args...)
}
