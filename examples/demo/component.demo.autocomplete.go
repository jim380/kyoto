package main

import (
	"strings"

	"github.com/kyoto-framework/kyoto"
)

type ComponentDemoAutocomplete struct {
	// Public
	Value       string
	Placeholder string
	Items       []string

	// Internal
	FilteredItems []string
}

func (c *ComponentDemoAutocomplete) Actions() kyoto.ActionMap {
	return kyoto.ActionMap{
		"Reload": func(args ...interface{}) {
			c.FilteredItems = []string{}
			if len(c.Value) != 0 {
				for _, item := range c.Items {
					if strings.Contains(
						strings.ToLower(item),
						strings.ToLower(c.Value),
					) {
						c.FilteredItems = append(c.FilteredItems, item)
					}
				}
			}
		},
		"Select": func(args ...interface{}) {
			c.Value = args[0].(string)
			c.FilteredItems = []string{}
		},
	}
}
