#!/bin/bash
# clean up the test cache
buildevents cmd $BUILDKITE_BUILD_ID $STEP_SPAN_ID go-clean -- go clean -testcache
if [[ ! $(buildevents cmd $BUILDKITE_BUILD_ID $STEP_SPAN_ID go-test -- gotestsum --format testname --junitfile unit-tests.xml --junitfile-testcase-classname relative -- -coverprofile=cover.out ./...) ]]; then
    buildevents step $BUILDKITE_BUILD_ID $STEP_SPAN_ID $STEP_START command
    exit 1
fi

buildevents step $BUILDKITE_BUILD_ID $STEP_SPAN_ID $STEP_START command
