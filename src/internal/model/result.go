package model

import "time"

type SuiteResult struct {
	Tag         string
	TestsResult []TestResult
}

type TestResult struct {
	HTTPResponse HTTPResponse
	Assertion    Assert
	Duration     time.Duration
}
