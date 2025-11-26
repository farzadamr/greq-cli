package config

import (
	"fmt"
	"os"

	"github.com/farzadamr/greq-cli/internal/model"
	"go.yaml.in/yaml/v3"
)

func Load(path string) (*model.SuiteFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var suite model.SuiteFile
	if err = yaml.Unmarshal(data, &suite); err != nil {
		return nil, fmt.Errorf("invalid yaml: %w", err)
	}

	err = validateSuite(&suite)
	if err != nil {
		return nil, err
	}

	applyDefaults(&suite)

	return &suite, nil
}
