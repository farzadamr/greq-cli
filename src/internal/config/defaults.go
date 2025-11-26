package config

import (
	"github.com/farzadamr/greq-cli/internal/model"
)

func applyDefaults(sf *model.SuiteFile) {
	if sf.Global.Timeout == 0 {
		sf.Global.Timeout = 5000
	}

	if sf.Vars == nil {
		sf.Vars = make(map[string]string)
	}

	for si := range sf.Suites {
		suite := &sf.Suites[si]

		for ti := range suite.Tests {
			test := &suite.Tests[ti]

			if test.Method == "" {
				test.Method = "GET"
			}

			if test.Save == nil {
				test.Save = make(map[string]string)
			}
		}
	}

}
