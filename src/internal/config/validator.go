package config

import (
	"errors"
	"fmt"

	"github.com/farzadamr/greq-cli/internal/model"
)

func validateSuite(sf *model.SuiteFile) error {
	if sf.Version == "" {
		return errors.New("version feild is required")
	}

	if len(sf.Env) == 0 {
		return errors.New("env section is required")
	}

	if len(sf.Suites) == 0 {
		return errors.New("no suite defined")
	}

	for _, suite := range sf.Suites {
		if suite.Tag == "" {
			return errors.New("suite tag is required")
		}

		if len(suite.Tests) == 0 {
			return fmt.Errorf("no test defined in suite with tag '%s'", suite.Tag)
		}

		for _, test := range suite.Tests {

			if test.Path == "" {
				return fmt.Errorf("path is required in test with tag '%s'", suite.Tag)
			}

			if test.Assert.Status < 100 {
				return fmt.Errorf("test in tag '%s' has invalid status assertion", suite.Tag)
			}

		}
	}
	return nil
}
