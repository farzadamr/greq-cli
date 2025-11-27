package model

import "time"

type SuiteResult struct {
	Tag            string
	TestsResponses []TestResponse
}

type TestResponse struct {
	HTTPResponse HTTPResponse
	Assertion    Assert
	Duration     time.Duration
}

type VerifyResult struct {
	Tag         string
	TestResults []TestResult
}

type TestResult struct {
	Path     string
	Method   string
	Duration time.Duration
	Assert   []AssertResult
}

type AssertResult struct {
	Title   string
	Message string
	Passed  bool
}
