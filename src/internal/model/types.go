package model

type SuiteFile struct {
	Version string
	Env     map[string]string
	Vars    map[string]string
	Global  GlobalConfig
	Suites  []Suite
}

type GlobalConfig struct {
	Timeout int64
	Headers map[string]string
}

type Suite struct {
	Tag   string
	Tests []Test
}

type Test struct {
	Method string
	Path   string
	Body   any
	Save   map[string]string
	Assert Assert
}

type Assert struct {
	Status   int
	Contains []string
}
