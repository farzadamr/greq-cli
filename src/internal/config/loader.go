package config

import (
	"fmt"
	"strings"

	"github.com/farzadamr/greq-cli/internal/model"
	"github.com/spf13/viper"
)

func LoadSuite(path string) (*model.TestSuite, error) {
	v := viper.New()
	v.SetConfigFile(path)

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var suite model.TestSuite
	if err := v.Unmarshal(&suite); err != nil {
		return nil, err
	}

	for i, _ := range suite.Tests {
		suite.Tests[i].Request.URL = replaceVars(suite.Tests[i].Request.URL, suite.Env)
	}
	return &suite, nil
}

func replaceVars(s string, env map[string]string) string {
	for k, v := range env {
		placeholder := fmt.Sprintf("{%s}", k)
		s = strings.ReplaceAll(s, placeholder, v)
	}
	return s
}
