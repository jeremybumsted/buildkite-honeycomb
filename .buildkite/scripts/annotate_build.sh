#!/bin/bash


echo "+++ Finishing up the build in Honeycomb"

trace_url=$(buildevents build "$BUILDKITE_BUILD_ID" $BUILD_START success)

annotation=$(cat << EOF
# Honeycomb Buildevents
The trace for this build can be found at $trace_url
EOF)

echo "~~~ Annotating build"
echo "$annotation" | buildkite-agent annotate --context "honeycomb" --style "info"
