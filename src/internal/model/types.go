package model

type TestSuite struct {
	Env   map[string]string `yaml:"env"`
	Tests []APITest         `yaml:"tests"`
}

type APITest struct {
	Name    string `yaml:"name"`
	Request HttpRequest
	Expect  Expectation `yaml:"expect"`
}

type Expectation struct {
	Status   int    `yaml:"status"`
	Contains string `yaml:"contains,omitempty"`
}
