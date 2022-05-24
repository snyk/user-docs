# Tool: snyk-filter

[![Snyk logo](https://snyk.io/style/asset/logo/snyk-print.svg)](https://snyk.io)

## Custom Filtering for Snyk CLI

snyk-filter takes the JSON outputted from the [Snyk CLI](https://support.snyk.io/hc/en-us/articles/360003812578-CLI-reference), e.g. `snyk test --json` and applies custom filtering of the results, as well as options to fail your build.

[![Known Vulnerabilities](https://snyk.io/test/github/snyk-tech-services/snyk-filter/badge.svg?targetFile=package.json)](https://snyk.io/test/github/snyk-tech-services/snyk-filter?targetFile=package.json) [![CircleCI](https://circleci.com/gh/snyk-tech-services/snyk-filter.svg?style=svg)](https://circleci.com/gh/snyk-tech-services/snyk-filter)

## How do I use it?

### Clone & Install

First, clone the repo. Then run

`npm install -g`

#### Note about `node-jq`

snyk-filter uses the `node-jq` library, which requires that a [`jq`](https://stedolan.github.io/jq/) binary is installed. This typically happens transparently via `npm install -g`, but on some systems JQ does not get properly installed locally. If you receive an error after installation regarding `node-jq`, then `jq` should be installed manually to avoid this error.

```
# install jq ahead of time (ubuntu example)
sudo apt-get install -y jq

# tell node-jq to skip trying to install it on its own
export NODE_JQ_SKIP_INSTALL_BINARY=true

# tell node-jq where the existing jq binary is
export JQ_PATH=$(which jq)

# finally, install snyk-filter (does not work with node version > 12)
sudo npm install -g
```

### Usage

1. Implement your custom JQ filters in a .snyk-filter/snyk.yml file relative to your current working directory where you will be running snyk test from (see in [sample-filters](https://github.com/snyk-tech-services/snyk-filter/tree/develop/sample-filters) and tweak things from there - use [JQPlay](https://jqplay.org/) )
2. Then pipe your `snyk test --json` output into `snyk-filter` or use the `-i` argument to input a json file. Use the `-f` argument to point to the yml file containing your custom filters if you are not using the default location (.snyk-filter/snyk.yml).
3. Return code of snyk-filter will be 0 for pass (no issues) and 1 for fail (issues found)

#### Example with Snyk CLI (using .snyk-filter/snyk.yml by default)

`snyk test --json | snyk-filter`

#### Example with Snyk CLI and custom yml file location

`snyk test --json | snyk-filter -f /path/to/example-cvss-9-or-above.yml`

#### Example

`snyk-filter -i snyk_results.json`

#### Example with custom yml file location

`snyk-filter -i snyk_results.json -f /path/to/example-high-upgradeable-vulns.yml`

### Options

`--json` to output json

#### License

License: Apache License, Version 2.0
