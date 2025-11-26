package model

import (
	"time"
)

type TestResult struct {
	Name     string
	Passed   bool
	Time     time.Duration
	Response HTTPResponse
	Expect   Expectation
	Err      *error
}
