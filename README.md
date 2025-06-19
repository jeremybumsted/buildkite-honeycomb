

# buildkite-honeycomb
An example repository for using Honeycomb's Buildevents and Buildkite Markers in a Buildkite pipeline.
[![Add to Buildkite](https://buildkite.com/button.svg)](https://buildkite.com/new)

## Setup

Fork the repository, then click the button below to create a new Pipeline.

Further setup instructions are below for each of the tools used in this repository.

### Honeycomb Buildevents
The first set of tests uses [Honeycomb Buildevents](https://github.com/honeycombio/buildevents).

In your agent's environment hook (the location depends on how you have [installed the agent](https://buildkite.com/docs/agent/v3/installation)), set up the Honeycomb API key and dataset name:

`~/.buildkite-agent/hooks/environment`
```
#!/bin/bash
export BUILDEVENT_APIKEY=your_api_key_here
# Dataset to use. Default is buildevents, for this example, we'll use "buildkite"
export BUILDEVENT_DATASET=buildkite
export BUILDEVENT_APIHOST="https://api.honeycomb.io/"
export BUILDEVENT_CIPROVIDER="BUILDKITE"
```
