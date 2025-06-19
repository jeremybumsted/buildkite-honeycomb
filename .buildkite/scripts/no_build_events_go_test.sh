#!/bin/bash
# clean up the test cache
go clean -testcache
go-test -- gotestsum --format testname --junitfile unit-tests.xml --junitfile-testcase-classname relative -- -coverprofile=cover.out ./...
